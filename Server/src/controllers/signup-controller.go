package controllers

import (
	models "Server/src/models"
	"encoding/json"
	"net/http"
	// "gorm.io/gorm"
	// "fmt"
	// "strconv"
	// "github.com/gorilla/mux"
	utils "Server/src/utils"
)

var NewSignUp models.SignUp

func StoreSignupDetails(w http.ResponseWriter, r *http.Request) {
	StoreSignupDetails := &models.SignUp{}
	utils.ParseBody(r, StoreSignupDetails)
	s := StoreSignupDetails.StoreSignupDetails()
	res, _ := json.Marshal(s)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
