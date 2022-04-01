package main

import (
	"go_mysql/db/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func migrations() {
	dsn := "root:321123@tcp(127.0.0.1:3306)/go_mysql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("No connection to database")
	}
	db.Migrator().CreateTable(&models.User{})
	db.Migrator().CreateTable(&models.Todo{})
}

func main() {
	migrations()
}
