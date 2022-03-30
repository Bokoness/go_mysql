package main

import (
	"fmt"
	Services "go_mysql/services"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Todos    []Todo
}

type Todo struct {
	gorm.Model
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  int64  `json:"userId"`
	User    User
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = Services.Hash(u.Password)
	return
}

func seeders(db *gorm.DB) {
	users := []User{
		{
			Username: "bokoness",
			Password: "321123",
		},
	}
	todos := []Todo{}
	for i := 0; i < 10; i++ {
		uname := fmt.Sprintf("name%d", i)
		users = append(users, User{
			Username: uname,
			Password: "321123",
		})
		var j int64 = 1
		for ; j < 5; j++ {
			tname := fmt.Sprintf("todo%d%d", i, j)
			todos = append(todos, Todo{
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
	db.Migrator().CreateTable(&User{})
	db.Migrator().CreateTable(&Todo{})
	seeders(db)
}

func main() {
	migrations()
}
