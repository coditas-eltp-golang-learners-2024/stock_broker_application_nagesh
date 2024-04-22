package handlers

import (
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/service"
	"Stock_broker_application/src/app/authentication/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ForgotPasswordHandler(fpSvc service.ForgotPasswordService) gin.HandlerFunc {
	return func(c *gin.Context) {

		var userInput models.ForgotPassword
		if err := c.ShouldBindJSON(&userInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		if err := utils.ValidateForgotPassword(userInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := fpSvc.ForgotPassword(&userInput); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing request"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
	}
}
