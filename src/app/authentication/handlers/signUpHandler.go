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

// SignUp godoc
// @Summary Sign up a new user
// @Description Sign up a new user in the system
// @Tags authentication
// @Accept json
// @Produce json
// @Param userInput body models.User true "User data to sign up"
// @Router /signup [post]
func SignUp(signUpService *service.SignUpService) gin.HandlerFunc {
	return func(c *gin.Context) {

		var userInput models.User

		if err := c.ShouldBindJSON(&userInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			c.Abort()
			return
		}

		if err := validate.Struct(userInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrValidationFailed, "validation_errors": err.Error()})
			c.Abort()
			return
		}

		if err := signUpService.SignUp(&userInput); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Customer signed up successfully"})
	}
}
