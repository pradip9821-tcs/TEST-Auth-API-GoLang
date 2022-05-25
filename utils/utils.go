package utils

import (
	"Auth_Rest_Api/modles"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	gotenv.Load()
}

func RespondWithError(Message string, w http.ResponseWriter, status int) {
	var error modles.Error
	error.Message = Message
	error.Status = 0
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(error)
}

func ResponseJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}

func GenerateToken(user modles.User) (string, error) {
	var err error

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":     "course",
		"email":   user.Email,
		"user_id": user.ID,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		log.Println(err)
	}
	return tokenString, err
}
