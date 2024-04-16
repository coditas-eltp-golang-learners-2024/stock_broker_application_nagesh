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

func authMiddleware(c *gin.Context) {
	// Check authentication
	// Example: Check for Authorization header and validate token
	token := c.GetHeader("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	// Continue to next handler if authenticated
	c.Next()
}

// @Summary Sign in user
// @Description Sign in user with email and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param input body models.UserSignIn true "User sign in details"
// @Success 200 {object} gin.H{"message": "Email and Password verified"}
// @Failure 400 {object} gin.H{"error": "Bad Request", "validation_errors": "Validation error details"}
// @Failure 502 {object} gin.H{"error": "Bad Gateway", "validation_errors": "Email or password verification failed"}
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
	c.JSON(http.StatusOK, gin.H{"Token": token})

}
