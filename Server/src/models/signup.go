package models

import (
	
)


type SignUp struct{
	ID        uint      `gorm:"primaryKey" json:"id"`
	FirstName string 	`gorm:"not null;" json:"firstname"`
	LastName  string 	`gorm:"not null;" json:"lastname"`
	Email     string	`gorm:"not null;" json:"email"`
	Password  uint		`gorm:"not null;" json:"password"`
	ConPassword  uint	`gorm:"not null;" json:"conpassword"`
}
