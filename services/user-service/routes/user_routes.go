package routes

import (
	"services/user-service/controllers"
	"services/user-service/repositories"
	"services/user-service/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	repo := repositories.NewUserRepository(db)
	service := services.NewUserService(repo)
	controller := controllers.NewUserController(service)

	// Home Page
	router.GET("/", controllers.HomePage)
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/register", controller.RegisterUser)
	}
}
