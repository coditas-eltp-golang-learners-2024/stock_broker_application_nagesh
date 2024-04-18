package service

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/repo"
	"Stock_broker_application/src/app/authentication/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func SignInService(user *models.UserSignIn) (string, error) {
	utils.Logger.Println("Singin service called")
	result := repo.VerifySignInCredentials(user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Email,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenString, _ := token.SignedString(constants.JWTSECRETKEY)
	if !result {
		utils.Logger.Println("Email or password verification failed")
		return "", constants.ErrEmailOrPasswordVerificationFailed
	}
	utils.Logger.Println("Signin successful")
	return tokenString, nil

}
