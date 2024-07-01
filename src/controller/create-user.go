package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/matheustrres/go-first-crud/src/config/logger"
	"github.com/matheustrres/go-first-crud/src/config/validation"
	"github.com/matheustrres/go-first-crud/src/controller/models/request"
	"github.com/matheustrres/go-first-crud/src/controller/models/response"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Initializing CreateUser controller",
		zap.String("journey", "createUser"),
	)

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err)

		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)

		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"),
	)

	response := response.UserResponse{
		ID:    "random-id",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	c.JSON(200, response)
}
