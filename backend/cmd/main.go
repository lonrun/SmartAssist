package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lonrun/SmartAssist/backend/internal/app"
	"github.com/lonrun/SmartAssist/backend/internal/handlers"
	"github.com/lonrun/SmartAssist/backend/internal/services"
	"github.com/lonrun/SmartAssist/backend/pkg/config"
	"github.com/lonrun/SmartAssist/backend/pkg/database"
	"github.com/lonrun/SmartAssist/backend/pkg/logger"
)

func main() {
	// Load configuration from file
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize logger
	logger, err := logger.NewLogger(cfg.Logger.Level)
	if err != nil {
		log.Fatalf("Error initialize logger: %v", err)
	}

	// Initialize database connection
	db, err := database.NewDatabase(cfg.Database.Name, cfg.Database.Url)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Initialize services
	pdfService := services.NewPDFService(db)
	ticketService := services.NewTicketService(db)
	deviceService := services.NewDeviceService(db)

	// Initialize application
	application := app.NewApp(pdfService, ticketService, deviceService)

	// Initialize router
	router := gin.Default()

	// Add handlers
	router.POST("/pdf", handlers.PDFHandler(application))
	router.POST("/ticket", handlers.TicketHandler(application))
	router.POST("/device", handlers.DeviceHandler(application))

	// Start HTTP server
	logger.Info("Starting HTTP server...")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Error starting HTTP server: %v", err)
	}
}
