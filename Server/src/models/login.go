package models

import (
	"gorm.io/gorm"
)

type LogIn struct{
	gorm.Model

	UserName string   `gorm:"not null;" json:"username"`
	Password string   `gorm:"not null;" json:"password"`
}


func (l *LogIn) StoreLogInDetails() *LogIn{
	db.Create(&l)
	return l
}