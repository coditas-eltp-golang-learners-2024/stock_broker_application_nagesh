//actual api method

package handlers

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func SignUp(c *gin.Context) {

	var userInput models.User

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := validate.Struct(userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrValidationFailed, "validation_errors": err.Error()})
		return
	}

	if err := service.SignUpService(&userInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Customer signed up successfully"})
}
