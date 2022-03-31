package main

import (
	"fmt"
	models "go_mysql/db/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func seeders(db *gorm.DB) {
	users := []models.User{
		{
			Username: "bokoness",
			Password: "321123",
		},
	}
	todos := []models.Todo{}
	for i := 0; i < 10; i++ {
		uname := fmt.Sprintf("name%d", i)
		users = append(users, models.User{
			Username: uname,
			Password: "321123",
		})
		var j int64 = 1
		for ; j < 5; j++ {
			tname := fmt.Sprintf("todo%d%d", i, j)
			todos = append(todos, models.Todo{
				Title:   tname,
				Content: "content",
				UserID:  j,
			})
		}
	}
	db.Create(&users)
	db.Create(&todos)
}

func migrations() {
	dsn := "root:321123@tcp(127.0.0.1:3306)/go_mysql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("No connection to database")
	}
	db.Migrator().CreateTable(&models.User{})
	db.Migrator().CreateTable(&models.Todo{})
	seeders(db)
}

func main() {
	migrations()
}
