// Package routes contains all routes of the application
package routes

import (
	userController "github.com/gbrayhan/microservices-go/infrastructure/rest/controllers/user"
	"github.com/gin-gonic/gin"
)

// UserRoutes is a function that contains all routes of the user
func UserRoutes(router *gin.RouterGroup, controller *userController.Controller) {
	userrouter := router.Group("/user")
	{
		userrouter.POST("/", controller.NewUser)
		userrouter.GET("/:id", controller.GetUsersByID)
		userrouter.GET("/", controller.GetAllUsers)
		userrouter.PUT("/:id", controller.UpdateUser)
		userrouter.DELETE("/:id", controller.DeleteUser)
	}
}
