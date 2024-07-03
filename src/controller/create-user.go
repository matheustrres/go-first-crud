package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheustrres/go-first-crud/src/config/logger"
	"github.com/matheustrres/go-first-crud/src/config/validation"
	"github.com/matheustrres/go-first-crud/src/controller/models/request"
	"github.com/matheustrres/go-first-crud/src/model"
	"github.com/matheustrres/go-first-crud/src/view"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
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

	domain := model.NewUserDomain(
		userRequest.Name,
		userRequest.Email,
		userRequest.Password,
		userRequest.Age,
	)

	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}
