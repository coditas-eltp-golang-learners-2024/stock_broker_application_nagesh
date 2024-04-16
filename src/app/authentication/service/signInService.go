package service

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/repo"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func SignInService(user *models.UserSignIn) (string, error) {
	result := repo.VerifySignInCredentials(user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := token.SignedString(constants.JWTSECRETKEY)
	if !result {
		return "", constants.ErrEmailOrPasswordVerificationFailed
	}
	return tokenString, nil

}
