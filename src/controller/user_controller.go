package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/matheustrres/go-first-crud/src/model/service"
)

func NewUserControllerInterface(
	service service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service,
	}
}

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	FindUserByID(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
