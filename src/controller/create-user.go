package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/matheustrres/go-first-crud/src/config/validation"
	"github.com/matheustrres/go-first-crud/src/controller/models/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		fmt.Printf("Error trying to marshal object, error=%s\n", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)

		return
	}

	fmt.Println(userRequest)
}
