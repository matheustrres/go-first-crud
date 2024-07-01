package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/matheustrres/go-first-crud/src/config/rest_errors"
	"github.com/matheustrres/go-first-crud/src/controller/models/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_errors.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, error=%s", err.Error()),
		)

		c.JSON(restErr.Code, restErr)

		return
	}

	fmt.Println(userRequest)
}
