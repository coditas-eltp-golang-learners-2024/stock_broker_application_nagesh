package router

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/handlers"
	"Stock_broker_application/src/app/authentication/repo"
	"Stock_broker_application/src/app/authentication/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	userRepo := repo.NewUserInstance(db)
	userAuthService := service.NewSignInService(userRepo)
	router.POST(constants.SignIn, handlers.SignInHandler(userAuthService))

	userRepoSignUp := repo.NewUserSignUpInstance(db)
	userSignUpService := service.NewSignUpService(userRepoSignUp)
	router.POST(constants.SignUp, handlers.SignUp(userSignUpService))

	changePasswordRepo := repo.NewChangePasswordInstance(db)
	changePasswordService := service.NewChangePasswordService(changePasswordRepo)
	router.POST(constants.ChangePassword, handlers.AuthMiddleware(), handlers.ChangePasswordHandler(changePasswordService))

	userValidateOtpRepo := repo.NewValidateOtpInstance(db)
	userValidateOtpRepository := service.NewValidateOtpService(userValidateOtpRepo)
	router.POST(constants.ValidateOtp, handlers.AuthMiddleware(), handlers.ValidateOtp(userValidateOtpRepository))

	return router

}
