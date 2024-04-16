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

// @Summary Sign up a new user
// @Description Sign up a new user in the system
// @Tags authentication
// @Accept json
// @Produce json
// @Param userInput body models.User true "User data to sign up"
// @Success 201 {object} gin.H{"message": "Customer signed up successfully"}
// @Failure 400 {object} gin.H{"error": "Bad Request", "details": "Validation failed"}
// @Failure 500 {object} gin.H{"error": "Internal Server Error", "details": "Failed to sign up user"}
// @Router /signup [post]

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
