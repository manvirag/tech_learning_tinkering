package handlers

import (
	"log"
	"net/http"

	"ot-editor/internal/models"
	"ot-editor/internal/services"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// WebSocketHandler handles WebSocket connections for collaborative editing
type WebSocketHandler struct {
	documentService *services.DocumentService
	upgrader        websocket.Upgrader
}

// NewWebSocketHandler creates a new WebSocket handler
func NewWebSocketHandler(documentService *services.DocumentService) *WebSocketHandler {
	return &WebSocketHandler{
		documentService: documentService,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				// Allow connections from any origin (adjust for production)
				return true
			},
		},
	}
}

// HandleWebSocket upgrades HTTP connections to WebSocket and handles real-time communication
func (wh *WebSocketHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to WebSocket
	conn, err := wh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}
	defer conn.Close()

	// Generate a unique client ID
	clientID := uuid.New().String()
	log.Printf("New WebSocket connection established for client %s", clientID)

	// Create a wrapper for the WebSocket connection
	clientConn := &WebSocketConnection{conn: conn}

	// Handle the WebSocket connection
	wh.handleClient(clientID, clientConn)
}

// WebSocketConnection wraps the gorilla/websocket connection
type WebSocketConnection struct {
	conn *websocket.Conn
}

// WriteMessage writes a message to the WebSocket connection
func (wsc *WebSocketConnection) WriteMessage(messageType int, data []byte) error {
	return wsc.conn.WriteMessage(messageType, data)
}

// Close closes the WebSocket connection
func (wsc *WebSocketConnection) Close() error {
	return wsc.conn.Close()
}

// handleClient handles communication with a single client
func (wh *WebSocketHandler) handleClient(clientID string, conn *WebSocketConnection) {
	var documentID string
	var joined bool

	// Set up ping/pong handlers for connection health
	conn.conn.SetPingHandler(func(string) error {
		return conn.WriteMessage(websocket.PongMessage, nil)
	})

	// Clean up when client disconnects
	defer func() {
		if joined && documentID != "" {
			wh.handleClientLeave(clientID, documentID)
		}
		log.Printf("Client %s disconnected", clientID)
	}()

	// Message handling loop
	for {
		// Read message from client
		_, messageData, err := conn.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error for client %s: %v", clientID, err)
			}
			break
		}

		// Parse the message
		message, err := models.MessageFromJSON(messageData)
		if err != nil {
			log.Printf("Failed to parse message from client %s: %v", clientID, err)
			wh.sendError(conn, "INVALID_MESSAGE", "Failed to parse message")
			continue
		}

		// Handle the message based on its type
		switch message.Type {
		case models.MsgJoinDocument:
			doc, _, err := wh.handleJoinDocument(clientID, message, conn)
			if err != nil {
				wh.sendError(conn, "JOIN_FAILED", err.Error())
				continue
			}
			documentID = doc.ID
			joined = true
			log.Printf("Client %s joined document %s", clientID, documentID)

		case models.MsgLeaveDocument:
			if joined {
				wh.handleLeaveDocument(clientID, message)
				joined = false
			}

		case models.MsgOperation:
			if !joined {
				wh.sendError(conn, "NOT_JOINED", "Must join a document first")
				continue
			}
			wh.handleOperation(clientID, message)

		case models.MsgCursorUpdate:
			if !joined {
				wh.sendError(conn, "NOT_JOINED", "Must join a document first")
				continue
			}
			wh.handleCursorUpdate(clientID, message)

		case models.MsgRequestState:
			if !joined {
				wh.sendError(conn, "NOT_JOINED", "Must join a document first")
				continue
			}
			wh.handleRequestState(clientID, documentID, conn)

		case models.MsgPing:
			wh.sendPong(conn)

		default:
			log.Printf("Unknown message type from client %s: %s", clientID, message.Type)
			wh.sendError(conn, "UNKNOWN_MESSAGE_TYPE", "Unknown message type")
		}
	}
}

// handleJoinDocument handles a client joining a document
func (wh *WebSocketHandler) handleJoinDocument(clientID string, message *models.WebSocketMessage, conn *WebSocketConnection) (*models.Document, *models.Client, error) {
	var payload models.JoinDocumentPayload
	if err := message.ParsePayload(&payload); err != nil {
		return nil, nil, err
	}

	// Join the document
	doc, client, err := wh.documentService.JoinDocument(clientID, payload.DocumentID, payload.ClientName)
	if err != nil {
		return nil, nil, err
	}

	// Set the WebSocket connection in the client object so broadcasting works
	client.Conn = conn.conn

	// Send document state to the joining client
	stateMessage := models.CreateDocumentStateMessage(doc, doc.GetAllClients())
	wh.sendMessage(conn, stateMessage)

	// Notify other clients that a new client joined
	joinedMessage := models.CreateClientJoinedMessage(payload.DocumentID, client)
	wh.documentService.BroadcastToDocument(payload.DocumentID, clientID, joinedMessage)

	return doc, client, nil
}

// handleLeaveDocument handles a client leaving a document
func (wh *WebSocketHandler) handleLeaveDocument(clientID string, message *models.WebSocketMessage) {
	var payload models.LeaveDocumentPayload
	if err := message.ParsePayload(&payload); err != nil {
		log.Printf("Failed to parse leave document payload: %v", err)
		return
	}

	wh.handleClientLeave(clientID, payload.DocumentID)
}

// handleClientLeave handles the actual client leave logic
func (wh *WebSocketHandler) handleClientLeave(clientID, documentID string) {
	// Leave the document
	err := wh.documentService.LeaveDocument(clientID, documentID)
	if err != nil {
		log.Printf("Failed to leave document: %v", err)
		return
	}

	// Notify other clients that the client left
	leftMessage := models.CreateClientLeftMessage(documentID, clientID)
	wh.documentService.BroadcastToDocument(documentID, clientID, leftMessage)
}

// handleOperation handles document operations with OT
func (wh *WebSocketHandler) handleOperation(clientID string, message *models.WebSocketMessage) {
	var payload models.OperationPayload
	if err := message.ParsePayload(&payload); err != nil {
		log.Printf("Failed to parse operation payload: %v", err)
		return
	}

	log.Printf("Client %s applying operation: %s to document %s", clientID, payload.Operation.String(), payload.DocumentID)

	// Apply the operation using OT
	transformedOp, err := wh.documentService.ApplyOperation(clientID, payload.DocumentID, &payload.Operation)
	if err != nil {
		log.Printf("Failed to apply operation: %v", err)

		// Send error acknowledgment
		ackMessage := models.CreateOperationAckMessage(
			payload.DocumentID, payload.Operation, false, err.Error(), 0)
		wh.documentService.SendToClient(clientID, ackMessage)
		return
	}

	// Get the updated document version
	doc, err := wh.documentService.GetDocumentState(payload.DocumentID)
	if err != nil {
		log.Printf("Failed to get document state: %v", err)
		return
	}

	log.Printf("Operation applied successfully, new version: %d", doc.Version)

	// Send acknowledgment to the original client
	ackMessage := models.CreateOperationAckMessage(
		payload.DocumentID, *transformedOp, true, "", doc.Version)
	wh.documentService.SendToClient(clientID, ackMessage)

	// Broadcast the transformed operation to other clients
	broadcastMessage := models.CreateOperationMessage("", payload.DocumentID, *transformedOp)
	broadcastMessage.Type = models.MsgOperationBcast

	log.Printf("Broadcasting operation to other clients in document %s", payload.DocumentID)
	wh.documentService.BroadcastToDocument(payload.DocumentID, clientID, broadcastMessage)
}

// handleCursorUpdate handles cursor position updates
func (wh *WebSocketHandler) handleCursorUpdate(clientID string, message *models.WebSocketMessage) {
	var payload models.CursorUpdatePayload
	if err := message.ParsePayload(&payload); err != nil {
		log.Printf("Failed to parse cursor update payload: %v", err)
		return
	}

	// Update cursor position
	err := wh.documentService.UpdateCursor(clientID, payload.DocumentID, payload.Position)
	if err != nil {
		log.Printf("Failed to update cursor: %v", err)
		return
	}

	// Broadcast cursor update to other clients
	cursorMessage := models.CreateCursorUpdateMessage(clientID, payload.DocumentID, payload.Position)
	cursorMessage.Type = models.MsgCursorBcast
	wh.documentService.BroadcastToDocument(payload.DocumentID, clientID, cursorMessage)
}

// handleRequestState handles requests for current document state
func (wh *WebSocketHandler) handleRequestState(clientID, documentID string, conn *WebSocketConnection) {
	doc, err := wh.documentService.GetDocumentState(documentID)
	if err != nil {
		wh.sendError(conn, "STATE_ERROR", err.Error())
		return
	}

	// Send current document state
	stateMessage := models.CreateDocumentStateMessage(doc, doc.GetAllClients())
	wh.sendMessage(conn, stateMessage)
}

// Helper methods

// sendMessage sends a message to a WebSocket connection
func (wh *WebSocketHandler) sendMessage(conn *WebSocketConnection, message *models.WebSocketMessage) {
	data, err := message.ToJSON()
	if err != nil {
		log.Printf("Failed to serialize message: %v", err)
		return
	}

	err = conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
	}
}

// sendError sends an error message to a WebSocket connection
func (wh *WebSocketHandler) sendError(conn *WebSocketConnection, code, message string) {
	errorMessage := models.CreateErrorMessage(code, message)
	wh.sendMessage(conn, errorMessage)
}

// sendPong sends a pong message in response to ping
func (wh *WebSocketHandler) sendPong(conn *WebSocketConnection) {
	pongMessage := models.NewWebSocketMessage(models.MsgPong, nil, "")
	wh.sendMessage(conn, pongMessage)
}
