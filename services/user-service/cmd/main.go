package main

import (
	"log"
	"services/user-service/config"
	"services/user-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize database connection
	db := config.ConnectDB()
	defer config.CloseDB()

	// Setup router
	r := gin.Default()
	routes.RegisterRoutes(r, db)

	// Start server
	port := config.GetEnv("PORT", "8080")
	log.Printf("User Service running on port %s\n", port)
	r.Run(":" + port)
}
