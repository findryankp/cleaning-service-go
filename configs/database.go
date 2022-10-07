package configs

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	USER := os.Getenv("DBUSERNAME")
	PASS := os.Getenv("DBPASS")
	HOST := os.Getenv("DBHOST")
	PORT := os.Getenv("DBPORT")
	DBNAME := os.Getenv("DBNAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
}
