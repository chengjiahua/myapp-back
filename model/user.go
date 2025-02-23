package model

import (
	"myapp-back/database"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func Migrate() {
	database.DB.AutoMigrate(&User{})
}