package controllers

import (
	"Auth_Rest_Api/driver"
	"Auth_Rest_Api/modles"
	"Auth_Rest_Api/utils"
	"database/sql"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

var db *sql.DB

func Register(w http.ResponseWriter, r *http.Request) {
	db = driver.ConnectDB()
	var user modles.User

	_ = json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		utils.RespondWithError(`Email can't be null`, w, http.StatusBadRequest)
		return
	}
	if user.Password == "" {
		utils.RespondWithError(`Password can't be null`, w, http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Println(err)
	}

	user.Password = string(hash)

	//insert, err := db.Query("INSERT INTO user (email, password)VALUES ( 'test1@gmail.com', '1234' )")
	insert, err := db.Query(`INSERT INTO user (email, password)VALUES ('` + user.Email + `','` + user.Password + `')`)

	if err != nil {
		utils.RespondWithError("Server error!", w, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	defer insert.Close()

	sqlStatement := `SELECT id FROM user WHERE email=?`
	row := db.QueryRow(sqlStatement, user.Email)
	err = row.Scan(&user.ID)
	if err != nil {
		utils.RespondWithError("Server error!", w, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	response := modles.Response{
		Message: "Signup Successfully.",
		Data: modles.Data{
			Id:    user.ID,
			Email: user.Email,
		},
		Status: 1,
	}
	utils.ResponseJSON(w, response)
}

func Protected(w http.ResponseWriter, r *http.Request) {
	response := modles.Response{
		Message: "You are in protected router",
		Status:  1,
	}
	utils.ResponseJSON(w, response)
}

func Login(w http.ResponseWriter, r *http.Request) {
	db = driver.ConnectDB()
	var user modles.User
	var jwt modles.JWT
	_ = json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		utils.RespondWithError(`Email can't be null`, w, http.StatusBadRequest)
		return
	}
	if user.Password == "" {
		utils.RespondWithError(`Password can't be null`, w, http.StatusBadRequest)
		return
	}

	password := user.Password

	sqlStatement := `SELECT * FROM user WHERE email=?`
	row := db.QueryRow(sqlStatement, user.Email)
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.RespondWithError("User not Exist!", w, http.StatusNotFound)
			log.Println(err)
			return
		}
		utils.RespondWithError("Server error!", w, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	hashedPassword := user.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		utils.RespondWithError("Invalid Password!", w, http.StatusUnauthorized)
		log.Println("Invalid Password!")
		return
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		utils.RespondWithError("Token Generation Failed!!", w, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	jwt.Token = token

	response := modles.Response{
		Message: "Login Successfully.",
		Data: modles.Data{
			Id:    user.ID,
			Email: user.Email,
			Token: token,
		},
		Status: 1,
	}

	utils.ResponseJSON(w, response)
}
