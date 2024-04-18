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

func init() {
	validate = validator.New()
}

func ChangePasswordHandler(c *gin.Context) {

	var userInput models.ChangePassword

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := validate.Struct(userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrValidationFailed, "validation_errors": err.Error()})
		return
	}

	result, err := service.ChangePasswordService(&userInput)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": constants.ErrEmailOrPasswordVerificationFailed, "validation_errors": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": result})

}
