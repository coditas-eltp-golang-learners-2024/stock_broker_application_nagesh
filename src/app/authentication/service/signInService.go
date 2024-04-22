// package service

// import (
// 	"Stock_broker_application/src/app/authentication/constants"
// 	"Stock_broker_application/src/app/authentication/models"
// 	"Stock_broker_application/src/app/authentication/repo"
// 	"Stock_broker_application/src/app/authentication/utils"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// )

// func SignInService(user *models.UserSignIn) (string, error) {
// 	utils.Logger.Println("Singin service called")
// 	result := repo.VerifySignInCredentials(user)
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"username": user.Email,
// 		"exp":      time.Now().Add(time.Hour * 1).Unix(),
// 	})
// 	tokenString, _ := token.SignedString(constants.JWTSECRETKEY)
// 	if !result {
// 		utils.Logger.Println("Email or password verification failed")
// 		return "", constants.ErrEmailOrPasswordVerificationFailed
// 	}
// 	utils.Logger.Println("Signin successful")
// 	return tokenString, nil

// }

package service

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/models"
	"Stock_broker_application/src/app/authentication/repo"
	"Stock_broker_application/src/app/authentication/utils"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// SignInService handles sign-in logic
type SignInService struct {
	UserRepository repo.UserSignInRepository
}

// NewSignInService creates a new instance of SignInService
func NewSignInService(userRepository repo.UserSignInRepository) *SignInService {
	return &SignInService{
		UserRepository: userRepository,
	}
}

// SignIn authenticates the user
func (service *SignInService) SignIn(signInRequest models.UserSignIn) error {
	utils.Logger.Println("Sign in service called")
	user, err := service.UserRepository.GetUserByEmail(signInRequest.Email)

	if err != nil {
		utils.Logger.Println("Error occurred while getting user by email: ", err)
		return err
	}

	if user == nil {
		utils.Logger.Println("User does not exist")
		return constants.ErrUserNotExists
	}

	utils.Logger.Println("User exists")
	if user.Password != signInRequest.Password {
		utils.Logger.Println("Invalid credentials")
		return constants.ErrInvalidCredentials
	}

	utils.Logger.Println("Authentication successful")
	// Authentication successful
	return nil
}

func (service *SignInService) GenerateJWTTokenService(signInRequest *models.UserSignIn) string {
	utils.Logger.Println("Singin service called")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": signInRequest.Email,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenString, _ := token.SignedString(constants.JWTSECRETKEY)
	utils.Logger.Println("Signin successful")
	return tokenString
}

func (service *SignInService) GenerateOtp(user *models.UserSignIn) error {
	randSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randSource)
	otp := random.Intn(10001)

	otpExpTime := CreateOtpExpTime(1)
	currentUser, err := service.UserRepository.GetUserByEmail(user.Email)
	utils.Logger.Println("Generated Otp", "otp", otp, "user_id", currentUser)
	otp_err := service.UserRepository.SetOtpAndOtpExpiry(currentUser, otp, otpExpTime)
	if otp_err != nil {
		utils.Logger.Println("Error while setting Otp:", otp_err)
		return otp_err
	}
	if err != nil {
		utils.Logger.Println("Error while generating Otp:", err)
		return err
	}
	return nil
}

func CreateOtpExpTime(hour int) time.Time {
	currentTimeUTC := time.Now().UTC()
	expTime := currentTimeUTC.Add(time.Hour * time.Duration(hour))
	utils.Logger.Println("Generated Otp Expiry Time:", expTime)
	return expTime
}
