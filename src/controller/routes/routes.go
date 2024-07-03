package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/matheustrres/go-first-crud/src/controller"
)

func InitRoutes(
	r *gin.RouterGroup,
	c controller.UserControllerInterface,
) {
	r.GET("/get-user-by-id/:id", c.FindUserByID)
	r.GET("/get-user-by-email/:email", c.FindUserByEmail)
	r.POST("/users", c.CreateUser)
	r.PUT("/users/:id", c.UpdateUser)
	r.DELETE("/users/:id", c.DeleteUser)
}
