package main

import (
	"log"
	"net/http"
	"os"

	"drive-mvp/internal/handlers"
	"drive-mvp/internal/services"

	"github.com/gorilla/mux"
)

const (
	StorageDir = "storage"
	FilesDir   = "storage/files"
	ChunksDir  = "storage/chunks"
	Port       = ":8080"
)

func main() {
	// Create storage directories
	if err := createDirectories(); err != nil {
		log.Fatalf("Error creating directories: %v", err)
	}

	// Initialize services
	chunkService := services.NewChunkService(ChunksDir)
	fileService := services.NewFileService(FilesDir, chunkService)

	// Initialize handlers
	handler := handlers.NewHandler(fileService)

	// Setup routes
	router := setupRoutes(handler)

	// Start server
	log.Printf("Server starting on port %s", Port)
	log.Printf("Open http://localhost%s to access the web interface", Port)

	if err := http.ListenAndServe(Port, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func createDirectories() error {
	dirs := []string{FilesDir, ChunksDir, "web/static"}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	return nil
}

func setupRoutes(handler *handlers.Handler) *mux.Router {
	router := mux.NewRouter()

	// Web interface
	router.HandleFunc("/", handler.HomeHandler).Methods("GET")

	// API routes
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/upload", handler.UploadHandler).Methods("POST")
	api.HandleFunc("/files", handler.ListFilesHandler).Methods("GET")
	api.HandleFunc("/files/{fileId}", handler.DownloadHandler).Methods("GET")
	api.HandleFunc("/files/{fileId}", handler.UpdateHandler).Methods("PUT")
	api.HandleFunc("/files/{fileId}", handler.DeleteHandler).Methods("DELETE")

	// Static file serving
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/static/"))))

	// Add CORS middleware
	router.Use(corsMiddleware)

	return router
}

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
