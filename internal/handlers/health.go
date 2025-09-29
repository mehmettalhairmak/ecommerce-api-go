package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"ecommerce-api-go/internal/models"
)

// HealthHandler handles health check requests
type HealthHandler struct{}

// NewHealthHandler creates a new health handler
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Health returns the health status of the service
func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Health check response with system information
	healthInfo := models.HealthResponse{
		Status:    "healthy",
		Message:   "Server is up and running!",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Version:   "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(healthInfo); err != nil {
		log.Printf("Health check encoding error: %v", err)
		http.Error(w, "Health check failed", http.StatusInternalServerError)
		return
	}

	log.Printf("Health check completed in %v", time.Since(start))
}
