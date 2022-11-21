package controllers

import (
	"encoding/json"
	"net/http"
	models "Server/src/models"
	"gorm.io/gorm"
)

var Database *gorm.DB

func StoreSignupDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","pkglication/json")
	var signup models.SignUp
	json.NewDecoder(r.Body).Decode(&signup)
	Database.Create(&signup)
	json.NewEncoder(w).Encode(signup)
}