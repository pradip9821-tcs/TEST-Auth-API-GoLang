package controllers

import (
	"Auth_Rest_Api/modles"
	"Auth_Rest_Api/utils"
	"net/http"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	response := modles.Response{
		Message: "Welcome to homepage",
		Status:  1,
	}
	utils.ResponseJSON(w, response)
}
