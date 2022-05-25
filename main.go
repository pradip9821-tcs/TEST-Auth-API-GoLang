package main

import (
	"Auth_Rest_Api/controllers"
	"Auth_Rest_Api/driver"
	"Auth_Rest_Api/middleware"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	_ "golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {

	db = driver.ConnectDB()

	router := mux.NewRouter()

	router.HandleFunc("/", controllers.Homepage).Methods("GET")
	router.HandleFunc("/signup", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/protected", middleware.Auth(controllers.Protected)).Methods("GET")

	log.Println("Server Listening on 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
