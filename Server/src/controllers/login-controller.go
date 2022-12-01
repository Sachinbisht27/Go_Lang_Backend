package controllers

import (
	"encoding/json"
	"net/http"
	models "Server/src/models"
	// "gorm.io/gorm"
	// "fmt"
	// "strconv"
	// "github.com/gorilla/mux"
	utils "Server/src/utils"
)


var NewLogIn  models.LogIn

func StoreLogInDetails(w http.ResponseWriter, r *http.Request) {
	StoreLogInDetails := &models.LogIn{}
	utils.ParseBody(r, StoreLogInDetails)
	s := StoreLogInDetails.StoreLogInDetails()
	res, _ := json.Marshal(s)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}