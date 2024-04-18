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

// SignIn godoc
// @Summary Sign in to the system
// @Description Sign in to the system with user credentials
// @Tags Authentication
// @Accept json
// @Produce json
// @Param input body models.UserSignIn true "User credentials for signing in"
// @Router /signin [post]
func SignIn(c *gin.Context) {

	var userInput models.UserSignIn

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := validate.Struct(userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrValidationFailed, "validation_errors": err.Error()})
		return
	}

	token, err := service.SignInService(&userInput)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": constants.ErrEmailOrPasswordVerificationFailed, "validation_errors": err.Error()})
		return
	}
	if err := service.GenerateOtpService(&userInput); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": constants.ErrEmailOrPasswordVerificationFailed, "validation_errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully singed in", "Token": token})

}
