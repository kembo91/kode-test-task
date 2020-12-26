package userauth

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("storing_secret_like_this_is_wrong")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {

}
