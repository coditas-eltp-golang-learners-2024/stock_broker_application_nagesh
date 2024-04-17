package router

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST(constants.SignUp, handlers.SignUp)
	router.POST(constants.SignIn, handlers.SignIn)
	router.POST(constants.ChangePassword, handlers.AuthMiddleware(), handlers.ChangePasswordHandler)

	router.POST(constants.ValidateOtp, handlers.AuthMiddleware(), handlers.ValidateOtp)
	return router

}
