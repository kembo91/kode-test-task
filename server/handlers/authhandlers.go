package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/kembo91/kode-test-task/server/utils"

	"github.com/kembo91/kode-test-task/server/database"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("storing_secret_like_this_is_wrong")

//Claims jwt claims struct
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//AuthenticationMiddleware checks if a request is from a logged in user
func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		switch err {
		case http.ErrNoCookie:
			w.WriteHeader(http.StatusUnauthorized)
			return
		case nil:
			break
		default:
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		tokenString := c.Value
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		switch err {
		case jwt.ErrSignatureInvalid:
			w.WriteHeader(http.StatusUnauthorized)
			return
		case nil:
			break
		default:
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

//SignupHandler handles signup requests. Creates a new user in a database and provides with a token
func SignupHandler(db *database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var u database.Credentials
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			utils.JSONError(w, err, http.StatusInternalServerError)
			return
		}
		err = db.InsertUser(u)
		if err != nil {
			utils.JSONError(w, err, http.StatusBadRequest)
			return
		}
		setJWT(w, u.Username)
	}
}

//SigninHandler handles signin requests. Checks username and password and provides with a token
func SigninHandler(db *database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var u database.Credentials
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			utils.JSONError(w, err, http.StatusInternalServerError)
			return
		}
		err = db.CheckUser(u)
		if err != nil {
			utils.JSONError(w, err, http.StatusUnauthorized)
			return
		}
		setJWT(w, u.Username)
	}
}

//LogoutHandler handles logout requests by providing an expired token as a response
func LogoutHandler(db *database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func setJWT(w http.ResponseWriter, u string) {
	expTime := time.Now().Add(10 * time.Minute)
	claims := &Claims{
		Username: u,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expTime,
	})
}