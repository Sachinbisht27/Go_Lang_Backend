package models

import (
	"gorm.io/gorm"
	"Server/src/config"
	"fmt"
)

var db *gorm.DB
type SignUp struct{
	gorm.Model
	
	FullName string 	`gorm:"not null;" json:"fullname"`
	Email     string	`gorm:"not null;" json:"email"`
	Password  string	`gorm:"not null;" json:"password"`
	HasAgreed bool		`gorm:"not null;" json:"hasAgreed"`
}


func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&SignUp{}, &LogIn{})
	fmt.Println("Successfully connected!")
}

func (s *SignUp) StoreSignupDetails() *SignUp{
	// db.NewRecord(s)
	db.Create(&s)
	return s
}