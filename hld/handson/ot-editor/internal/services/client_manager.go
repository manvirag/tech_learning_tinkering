package services

import (
	"encoding/json"
	"log"
	"ot-editor/internal/models"
	"sync"
	"time"
)

// ClientManager manages WebSocket client connections
type ClientManager struct {
	clients map[string]*models.Client
	mutex   sync.RWMutex
}

// NewClientManager creates a new client manager
func NewClientManager() *ClientManager {
	return &ClientManager{
		clients: make(map[string]*models.Client),
	}
}

// AddClient adds a new client
func (cm *ClientManager) AddClient(client *models.Client) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	cm.clients[client.ID] = client
	log.Printf("Client %s (%s) connected", client.ID, client.Username)
}

// RemoveClient removes a client
func (cm *ClientManager) RemoveClient(clientID string) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	if client, exists := cm.clients[clientID]; exists {
		// Close WebSocket connection
		if client.Conn != nil {
			client.Conn.Close()
		}
		delete(cm.clients, clientID)
		log.Printf("Client %s removed", clientID)
	}
}

// GetClient gets a client by ID
func (cm *ClientManager) GetClient(clientID string) (*models.Client, bool) {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	client, exists := cm.clients[clientID]
	return client, exists
}

// GetAllClients returns all connected clients
func (cm *ClientManager) GetAllClients() []*models.Client {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	clients := make([]*models.Client, 0, len(cm.clients))
	for _, client := range cm.clients {
		clients = append(clients, client)
	}
	return clients
}

// SendToClient sends a message to a specific client
func (cm *ClientManager) SendToClient(clientID string, message *models.WebSocketMessage) error {
	cm.mutex.RLock()
	client, exists := cm.clients[clientID]
	cm.mutex.RUnlock()

	if !exists {
		log.Printf("Client %s not found for message delivery", clientID)
		return nil // Client not found, ignore silently
	}

	if client.Conn == nil {
		log.Printf("Client %s has no WebSocket connection", clientID)
		return nil // No connection available
	}

	// Update client's last seen time
	client.UpdateLastSeen()

	// Send message
	log.Printf("Sending message type '%s' to client %s", message.Type, clientID)
	err := client.Conn.WriteJSON(message)
	if err != nil {
		log.Printf("Failed to send message to client %s: %v", clientID, err)
		return err
	}

	log.Printf("Message sent successfully to client %s", clientID)
	return nil
}

// BroadcastToAll sends a message to all connected clients
func (cm *ClientManager) BroadcastToAll(message *models.WebSocketMessage) {
	cm.mutex.RLock()
	clients := make([]*models.Client, 0, len(cm.clients))
	for _, client := range cm.clients {
		clients = append(clients, client)
	}
	cm.mutex.RUnlock()

	for _, client := range clients {
		if err := cm.SendToClient(client.ID, message); err != nil {
			log.Printf("Error sending message to client %s: %v", client.ID, err)
			// Remove client if connection is broken
			cm.RemoveClient(client.ID)
		}
	}
}

// BroadcastToAllExcept sends a message to all clients except the specified one
func (cm *ClientManager) BroadcastToAllExcept(message *models.WebSocketMessage, exceptClientID string) {
	cm.mutex.RLock()
	clients := make([]*models.Client, 0, len(cm.clients))
	for _, client := range cm.clients {
		if client.ID != exceptClientID {
			clients = append(clients, client)
		}
	}
	cm.mutex.RUnlock()

	for _, client := range clients {
		if err := cm.SendToClient(client.ID, message); err != nil {
			log.Printf("Error sending message to client %s: %v", client.ID, err)
			// Remove client if connection is broken
			cm.RemoveClient(client.ID)
		}
	}
}

// CleanupInactiveClients removes clients that haven't been seen for a while
func (cm *ClientManager) CleanupInactiveClients(timeout time.Duration) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	now := time.Now()
	var toRemove []string

	for clientID, client := range cm.clients {
		if now.Sub(client.GetLastSeen()) > timeout {
			toRemove = append(toRemove, clientID)
		}
	}

	for _, clientID := range toRemove {
		if client, exists := cm.clients[clientID]; exists {
			if client.Conn != nil {
				client.Conn.Close()
			}
			delete(cm.clients, clientID)
			log.Printf("Cleaned up inactive client %s", clientID)
		}
	}
}

// GetClientCount returns the number of connected clients
func (cm *ClientManager) GetClientCount() int {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	return len(cm.clients)
}

// StartCleanupRoutine starts a goroutine that periodically cleans up inactive clients
func (cm *ClientManager) StartCleanupRoutine(interval, timeout time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			cm.CleanupInactiveClients(timeout)
		}
	}()
}

// HandleClientMessage processes incoming WebSocket messages
func (cm *ClientManager) HandleClientMessage(clientID string, messageData []byte, messageHandler func(*models.Client, *models.WebSocketMessage)) {
	client, exists := cm.GetClient(clientID)
	if !exists {
		log.Printf("Received message from unknown client %s", clientID)
		return
	}

	var message models.WebSocketMessage
	if err := json.Unmarshal(messageData, &message); err != nil {
		log.Printf("Error unmarshaling message from client %s: %v", clientID, err)
		return
	}

	// Update client's last seen time
	client.UpdateLastSeen()

	// Set client ID in message if not set
	if message.ClientID == "" {
		message.ClientID = clientID
	}

	// Call the message handler
	messageHandler(client, &message)
}

// PingClient sends a ping message to a client
func (cm *ClientManager) PingClient(clientID string) error {
	message := models.NewWebSocketMessage(models.MsgPing, nil, "")
	return cm.SendToClient(clientID, message)
}

// PingAllClients sends ping messages to all clients
func (cm *ClientManager) PingAllClients() {
	message := models.NewWebSocketMessage(models.MsgPing, nil, "")
	cm.BroadcastToAll(message)
}

// GetClientStats returns statistics about connected clients
func (cm *ClientManager) GetClientStats() map[string]interface{} {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	stats := map[string]interface{}{
		"total_clients": len(cm.clients),
		"clients":       make([]map[string]interface{}, 0, len(cm.clients)),
	}

	for _, client := range cm.clients {
		clientStats := map[string]interface{}{
			"id":        client.ID,
			"username":  client.Username,
			"last_seen": client.GetLastSeen(),
		}
		stats["clients"] = append(stats["clients"].([]map[string]interface{}), clientStats)
	}

	return stats
}
