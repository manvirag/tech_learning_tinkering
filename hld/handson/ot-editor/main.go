package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"ot-editor/internal/handlers"
	"ot-editor/internal/services"

	"github.com/gorilla/mux"
)

const (
	DefaultPort = "8080"
)

var documentService *services.DocumentService

func main() {
	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = DefaultPort
	}

	// Create storage directories
	if err := createDirectories(); err != nil {
		log.Fatalf("Error creating directories: %v", err)
	}

	// Initialize services with storage
	documentService = services.NewDocumentServiceWithStorage()

	// Setup graceful shutdown
	setupGracefulShutdown()

	// Initialize handlers
	httpHandler := handlers.NewHTTPHandler(documentService)
	wsHandler := handlers.NewWebSocketHandler(documentService)

	// Setup routes
	router := setupRoutes(httpHandler, wsHandler)

	// Start server
	log.Printf("🚀 Collaborative Document Editor (OT) starting on port %s", port)
	log.Printf("📝 Open http://localhost:%s to access the OT editor", port)
	log.Printf("🔌 WebSocket endpoint: ws://localhost:%s/ws", port)
	log.Printf("⚙️  Using Operational Transform (OT) for conflict resolution")
	log.Printf("💾 Documents auto-saved to storage/ot-documents/")
	log.Printf("🔄 Backups stored in storage/ot-backups/")

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func createDirectories() error {
	dirs := []string{
		"web/static",
		"storage/ot-documents",
		"storage/ot-backups",
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	log.Printf("Created storage directories for persistent document storage")
	return nil
}

func setupGracefulShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("\n🛑 Received shutdown signal, gracefully shutting down...")

		// Shutdown document service (saves all documents)
		if documentService != nil {
			documentService.Shutdown()
		}

		log.Println("✅ OT server shutdown complete")
		os.Exit(0)
	}()
}

func setupRoutes(httpHandler *handlers.HTTPHandler, wsHandler *handlers.WebSocketHandler) *mux.Router {
	router := mux.NewRouter()

	// WebSocket endpoint for real-time collaboration
	router.HandleFunc("/ws", wsHandler.HandleWebSocket)

	// Web interface routes
	router.HandleFunc("/", httpHandler.HomeHandler).Methods("GET")
	router.HandleFunc("/document/{documentId}", httpHandler.DocumentHandler).Methods("GET")

	// API routes for document management
	api := router.PathPrefix("/api").Subrouter()

	// Document CRUD operations
	api.HandleFunc("/documents", httpHandler.ListDocumentsHandler).Methods("GET")
	api.HandleFunc("/documents", httpHandler.CreateDocumentHandler).Methods("POST")
	api.HandleFunc("/documents/{documentId}", httpHandler.GetDocumentHandler).Methods("GET")
	api.HandleFunc("/documents/{documentId}", httpHandler.DeleteDocumentHandler).Methods("DELETE")

	// Document specific operations
	api.HandleFunc("/documents/{documentId}/history", httpHandler.GetDocumentHistoryHandler).Methods("GET")
	api.HandleFunc("/documents/{documentId}/clients", httpHandler.GetDocumentClientsHandler).Methods("GET")

	// Service health and statistics
	api.HandleFunc("/health", httpHandler.HealthHandler).Methods("GET")
	api.HandleFunc("/stats", httpHandler.StatsHandler).Methods("GET")

	// Static file serving
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/static/"))))

	// Add CORS middleware
	router.Use(corsMiddleware)

	// Add logging middleware
	router.Use(loggingMiddleware)

	return router
}

// CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Logging middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
