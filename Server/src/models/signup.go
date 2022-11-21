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

// var db *gorm.DB

// type Config struct {
// 	DB     *DBConfig
// }

// type DBConfig struct {
// 	Host     string
// 	Port     string
// 	Dbname   string
// 	Username string
// 	Password string
// }

// func GetServerConfig() *Config {
// 	loadEnvError := godotenv.Load(".env")
// 	if loadEnvError != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	return &Config{
// 		DB: &DBConfig{
// 			Port:     os.Getenv("PORT"),
// 			Host:     os.Getenv("HOST"),
// 			Dbname:   os.Getenv("DATABASE"),
// 			Username: os.Getenv("USERNAME"),
// 			Password: os.Getenv("PASSWORD"),
// 		},
// 	}
// }