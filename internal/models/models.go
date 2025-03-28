package models

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Claims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}
