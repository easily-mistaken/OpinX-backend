package main

import (
	"log"
	"os"

	"github.com/easily-mistaken/OpinX-backend/config"
	"github.com/easily-mistaken/OpinX-backend/routes"
	"github.com/easily-mistaken/OpinX-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	// Global state variables
	INRBalances   types.INRBalances
	OrderBook     types.OrderBook
	StockBalances types.StockBalances
)

func initializeGlobalState() {
	// Initialize with default values from config
	INRBalances = config.InitialINRBalances
	OrderBook = config.InitialOrderBook
	StockBalances = config.InitialStockBalances
}

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Error loading .env file")
	}

	// Initialize global state
	initializeGlobalState()

	// Get server port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	// Set Gin mode based on environment
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize Gin router
	r := gin.Default()

	// Basic health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"version": "1.0.0",
		})
	})

	// Initialize routes
	routes.Run(r)

	// Start server
	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	// Note: This line will never be reached as r.Run() blocks
	// fmt.Println("Server is running on port", port)
}
