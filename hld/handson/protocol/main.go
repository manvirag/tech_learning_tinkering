package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

// Message represents a simple message structure
type Message struct {
	ID      int       `json:"id"`
	Content string    `json:"content"`
	Time    time.Time `json:"time"`
}

// Global variables to store messages and clients
var (
	messages     []Message
	messagesLock sync.RWMutex
	clients      = make(map[chan Message]bool)
	clientsLock  sync.RWMutex
)

// addMessage adds a new message and notifies all clients
func addMessage(content string) {
	messagesLock.Lock()
	defer messagesLock.Unlock()

	msg := Message{
		ID:      len(messages) + 1,
		Content: content,
		Time:    time.Now(),
	}
	messages = append(messages, msg)

	// Notify all SSE clients
	clientsLock.RLock()
	for client := range clients {
		client <- msg
	}
	clientsLock.RUnlock()
}

// SSEHandler handles Server-Sent Events
func SSEHandler(w http.ResponseWriter, r *http.Request) {
	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Create a channel for this client
	messageChan := make(chan Message)
	clientsLock.Lock()
	clients[messageChan] = true
	clientsLock.Unlock()

	// Remove client when they disconnect
	defer func() {
		clientsLock.Lock()
		delete(clients, messageChan)
		clientsLock.Unlock()
		close(messageChan)
	}()

	// Send initial messages
	messagesLock.RLock()
	for _, msg := range messages {
		data, _ := json.Marshal(msg)
		fmt.Fprintf(w, "data: %s\n\n", data)
		w.(http.Flusher).Flush()
	}
	messagesLock.RUnlock()

	// Keep connection alive and send new messages
	for {
		select {
		case msg := <-messageChan:
			data, _ := json.Marshal(msg)
			fmt.Fprintf(w, "data: %s\n\n", data)
			w.(http.Flusher).Flush()
		case <-r.Context().Done():
			return
		}
	}
}

// LongPollHandler handles long polling requests
// its kind of same as https, it just waits for new messages
func LongPollHandler(w http.ResponseWriter, r *http.Request) {
	// Get the last message ID from query parameter
	lastID := 0
	if id := r.URL.Query().Get("last_id"); id != "" {
		fmt.Sscanf(id, "%d", &lastID)
	}

	// Wait for new messages
	for {
		messagesLock.RLock()
		if len(messages) > lastID {
			// New messages available
			newMessages := messages[lastID:]
			messagesLock.RUnlock()

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(newMessages)
			return
		}
		messagesLock.RUnlock()

		// Check if client disconnected
		select {
		case <-r.Context().Done():
			return
		case <-time.After(30 * time.Second):
			// Timeout after 30 seconds
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode([]Message{})
			return
		}
	}
}

// AddMessageHandler handles adding new messages
func AddMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	content := r.URL.Query().Get("content")
	if content == "" {
		http.Error(w, "Content is required", http.StatusBadRequest)
		return
	}

	addMessage(content)
	w.WriteHeader(http.StatusOK)
}

func main() {
	r := mux.NewRouter()

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// API endpoints
	r.HandleFunc("/sse", SSEHandler)
	r.HandleFunc("/longpoll", LongPollHandler)
	r.HandleFunc("/add", AddMessageHandler)

	// Start the server
	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
