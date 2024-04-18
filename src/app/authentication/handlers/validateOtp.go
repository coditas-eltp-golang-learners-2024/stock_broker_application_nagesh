//actual api method

package handlers

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Validate OTP for login
// @Description Validate OTP for login
// @Tags Authentication
// @Accept json
// @Produce json
// @Param input body models.ValidateOtp true "User OTP for login"
// @Router /validateotp [post]
func ValidateOtp(c *gin.Context) {

	var userInput models.ValidateOtp

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := validate.Struct(userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrValidationFailed, "validation_errors": err.Error()})
		return
	}

	if err := service.ValidateOtpService(userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP validated successfully"})
}
