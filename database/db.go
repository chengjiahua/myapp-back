package database

import (
	"myapp-back/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("mysql", config.Cfg.Database.Mysql.Dsn)
	if err != nil {
		panic("failed to connect database")
	}

	err = DB.DB().Ping()
	if err != nil {
		panic("failed to ping database")
	}

}

