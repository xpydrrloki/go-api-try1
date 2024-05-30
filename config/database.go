package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"go-api/helper"
	// "practice-api-gin-one/helper"
)

const (
	host = "localhost"
	port = 3306
	user = "root"
	password = "*p6Drr"
	dbName = "rest_go"

)

func DatabaseConnection() *gorm.DB {

	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := gorm.Open(mysql.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}