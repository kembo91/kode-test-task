package userauth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/kembo91/kode-test-task/server/database"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("storing_secret_like_this_is_wrong")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

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

func SignupHandler(db *database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var u database.Credentials
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = db.InsertUser(u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		setJWT(w, u.Username)
	}
}

func SigninHandler(db *database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var u database.Credentials
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		err = db.CheckUser(u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
		setJWT(w, u.Username)
	}
}

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
