package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/matheustrres/go-first-crud/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/get-user-by-id/:id", controller.FindUserByID)
	r.GET("/get-user-by-email/:email", controller.FindUserByEmail)
	r.POST("/users", controller.CreateUser)
	r.PUT("/users/:id", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)
}
