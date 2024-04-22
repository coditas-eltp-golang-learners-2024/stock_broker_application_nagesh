package constants

import "github.com/dgrijalva/jwt-go"

const (
	DatabaseYamlFilePath = "../../resources/application.yml"
	DatabasePost         = "3306"
	SuccessSignIn        = "Successfully Signed In"
)

var JWTSECRETKEY = []byte("aslkfdjadslkfjalsdjkf")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
