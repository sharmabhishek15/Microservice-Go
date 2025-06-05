package main

import (
	"log"

	"services/order-service/config"
	"services/order-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	defer config.CloseDB()

	router := gin.Default()
	routes.RegisterOrderRoutes(router, db)

	log.Println("Order Service is running on port 8082...")
	router.Run(":8082")
}
