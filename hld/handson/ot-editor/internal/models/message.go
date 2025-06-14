package models

import (
	"encoding/json"
	"time"
)

// MessageType defines the type of WebSocket message
type MessageType string

const (
	// Client to Server messages
	MsgJoinDocument  MessageType = "join_document"
	MsgLeaveDocument MessageType = "leave_document"
	MsgOperation     MessageType = "operation"
	MsgCursorUpdate  MessageType = "cursor_update"
	MsgRequestState  MessageType = "request_state"

	// Server to Client messages
	MsgDocumentState  MessageType = "document_state"
	MsgOperationAck   MessageType = "operation_ack"
	MsgOperationBcast MessageType = "operation_broadcast"
	MsgClientJoined   MessageType = "client_joined"
	MsgClientLeft     MessageType = "client_left"
	MsgCursorBcast    MessageType = "cursor_broadcast"
	MsgError          MessageType = "error"
	MsgPing           MessageType = "ping"
	MsgPong           MessageType = "pong"
)

// WebSocketMessage represents a message sent over WebSocket
type WebSocketMessage struct {
	Type      MessageType `json:"type"`
	Payload   interface{} `json:"payload,omitempty"`
	ClientID  string      `json:"client_id,omitempty"`
	Timestamp int64       `json:"timestamp"`
}

// JoinDocumentPayload for joining a document
type JoinDocumentPayload struct {
	DocumentID string `json:"document_id"`
	ClientName string `json:"client_name"`
}

// LeaveDocumentPayload for leaving a document
type LeaveDocumentPayload struct {
	DocumentID string `json:"document_id"`
}

// OperationPayload for sending operations
type OperationPayload struct {
	DocumentID string    `json:"document_id"`
	Operation  Operation `json:"operation"`
}

// CursorUpdatePayload for cursor position updates
type CursorUpdatePayload struct {
	DocumentID string `json:"document_id"`
	Position   int    `json:"position"`
}

// DocumentStatePayload for sending document state
type DocumentStatePayload struct {
	Document *Document `json:"document"`
	Clients  []*Client `json:"clients"`
}

// OperationAckPayload for acknowledging operations
type OperationAckPayload struct {
	DocumentID   string    `json:"document_id"`
	Operation    Operation `json:"operation"`
	Success      bool      `json:"success"`
	ErrorMessage string    `json:"error_message,omitempty"`
	NewVersion   int       `json:"new_version"`
}

// ClientJoinedPayload for notifying client joins
type ClientJoinedPayload struct {
	DocumentID string  `json:"document_id"`
	Client     *Client `json:"client"`
}

// ClientLeftPayload for notifying client leaves
type ClientLeftPayload struct {
	DocumentID string `json:"document_id"`
	ClientID   string `json:"client_id"`
}

// ErrorPayload for error messages
type ErrorPayload struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// NewWebSocketMessage creates a new WebSocket message
func NewWebSocketMessage(msgType MessageType, payload interface{}, clientID string) *WebSocketMessage {
	return &WebSocketMessage{
		Type:      msgType,
		Payload:   payload,
		ClientID:  clientID,
		Timestamp: time.Now().UnixNano(),
	}
}

// ToJSON converts the message to JSON
func (msg *WebSocketMessage) ToJSON() ([]byte, error) {
	return json.Marshal(msg)
}

// FromJSON creates a message from JSON
func MessageFromJSON(data []byte) (*WebSocketMessage, error) {
	var msg WebSocketMessage
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

// ParsePayload parses the payload into a specific type
func (msg *WebSocketMessage) ParsePayload(target interface{}) error {
	payloadBytes, err := json.Marshal(msg.Payload)
	if err != nil {
		return err
	}
	return json.Unmarshal(payloadBytes, target)
}

// CreateJoinDocumentMessage creates a join document message
func CreateJoinDocumentMessage(clientID, documentID, clientName string) *WebSocketMessage {
	payload := JoinDocumentPayload{
		DocumentID: documentID,
		ClientName: clientName,
	}
	return NewWebSocketMessage(MsgJoinDocument, payload, clientID)
}

// CreateLeaveDocumentMessage creates a leave document message
func CreateLeaveDocumentMessage(clientID, documentID string) *WebSocketMessage {
	payload := LeaveDocumentPayload{
		DocumentID: documentID,
	}
	return NewWebSocketMessage(MsgLeaveDocument, payload, clientID)
}

// CreateOperationMessage creates an operation message
func CreateOperationMessage(clientID, documentID string, operation Operation) *WebSocketMessage {
	payload := OperationPayload{
		DocumentID: documentID,
		Operation:  operation,
	}
	return NewWebSocketMessage(MsgOperation, payload, clientID)
}

// CreateCursorUpdateMessage creates a cursor update message
func CreateCursorUpdateMessage(clientID, documentID string, position int) *WebSocketMessage {
	payload := CursorUpdatePayload{
		DocumentID: documentID,
		Position:   position,
	}
	return NewWebSocketMessage(MsgCursorUpdate, payload, clientID)
}

// CreateDocumentStateMessage creates a document state message
func CreateDocumentStateMessage(document *Document, clients []*Client) *WebSocketMessage {
	payload := DocumentStatePayload{
		Document: document,
		Clients:  clients,
	}
	return NewWebSocketMessage(MsgDocumentState, payload, "")
}

// CreateOperationAckMessage creates an operation acknowledgment message
func CreateOperationAckMessage(documentID string, operation Operation, success bool, errorMsg string, newVersion int) *WebSocketMessage {
	payload := OperationAckPayload{
		DocumentID:   documentID,
		Operation:    operation,
		Success:      success,
		ErrorMessage: errorMsg,
		NewVersion:   newVersion,
	}
	return NewWebSocketMessage(MsgOperationAck, payload, "")
}

// CreateClientJoinedMessage creates a client joined message
func CreateClientJoinedMessage(documentID string, client *Client) *WebSocketMessage {
	payload := ClientJoinedPayload{
		DocumentID: documentID,
		Client:     client,
	}
	return NewWebSocketMessage(MsgClientJoined, payload, "")
}

// CreateClientLeftMessage creates a client left message
func CreateClientLeftMessage(documentID, clientID string) *WebSocketMessage {
	payload := ClientLeftPayload{
		DocumentID: documentID,
		ClientID:   clientID,
	}
	return NewWebSocketMessage(MsgClientLeft, payload, "")
}

// CreateErrorMessage creates an error message
func CreateErrorMessage(code, message string) *WebSocketMessage {
	payload := ErrorPayload{
		Code:    code,
		Message: message,
	}
	return NewWebSocketMessage(MsgError, payload, "")
}
