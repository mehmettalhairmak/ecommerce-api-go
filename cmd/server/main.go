package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"github.com/stripe/stripe-go/v82"
	"golang.org/x/time/rate"

	"ecommerce-api-go/internal/config"
	"ecommerce-api-go/internal/handlers"
	"ecommerce-api-go/internal/middleware"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize Stripe
	stripe.Key = cfg.StripeSecretKey

	// Initialize handlers
	paymentHandler := handlers.NewPaymentHandler()
	healthHandler := handlers.NewHealthHandler()

	// Initialize middleware
	rateLimiter := middleware.NewRateLimiter(rate.Limit(10), 100) // 10 requests/second, burst of 100

	// Setup routes with middleware
	http.HandleFunc("/create-payment-intent",
		middleware.CORSMiddleware(
			middleware.LoggingMiddleware(
				rateLimiter.Middleware(paymentHandler.CreatePaymentIntent))))

	http.HandleFunc("/health",
		middleware.LoggingMiddleware(healthHandler.Health))

	// Start server
	if cfg.IsProduction {
		startProductionServer(cfg)
	} else {
		startDevelopmentServer(cfg)
	}
}

// startDevelopmentServer starts the server in development mode (HTTP)
func startDevelopmentServer(cfg *config.Config) {
	log.Printf("Server starting on port %s (HTTP - Development Mode)", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatal(err)
	}
}

// startProductionServer starts the server in production mode (HTTPS)
func startProductionServer(cfg *config.Config) {
	server := &http.Server{
		Addr:         ":443",
		Handler:      nil,
		TLSConfig:    &tls.Config{MinVersion: tls.VersionTLS12},
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Starting production server on :443 (HTTPS)")
	log.Fatal(server.ListenAndServeTLS(cfg.TLSCertPath, cfg.TLSKeyPath))
}
