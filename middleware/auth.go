package middleware

import (
	"Auth_Rest_Api/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"
	"strings"
)

func init() {
	gotenv.Load()
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")[1]

		token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("token is not valid")
			}
			return []byte(os.Getenv("SECRET")), nil
		})

		if err != nil {
			utils.RespondWithError("Invalid Token, Unauthorized!", w, http.StatusUnauthorized)
			log.Println(err)
			return
		}

		if token.Valid {
			next.ServeHTTP(w, r)
		} else {
			utils.RespondWithError("Invalid Token, Unauthorized!", w, http.StatusUnauthorized)
			log.Println(err)
			return
		}
	})
}
