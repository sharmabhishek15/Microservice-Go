package routes

import (
	"services/order-service/controllers"
	"services/order-service/repositories"
	"services/order-service/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterOrderRoutes(router *gin.Engine, db *gorm.DB) {
	repo := repositories.NewOrderRepository(db)
	service := services.NewOrderService(repo)
	controller := controllers.NewOrderController(service)

	orderRoutes := router.Group("/orders")
	{
		orderRoutes.POST("/", controller.CreateOrder)
		orderRoutes.GET("/:user_id", controller.GetOrdersByUserID)
	}
}
