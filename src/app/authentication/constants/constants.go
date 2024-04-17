package constants

import "github.com/dgrijalva/jwt-go"

const (
	DatabaseYamlFilePath = "../../resources/application.yml"
	DatabasePost         = "3306"
)

var JWTSECRETKEY = []byte("aslkfdjadslkfjalsdjkf")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
