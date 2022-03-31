package models

import (
	"go_mysql/db"
	"go_mysql/services"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Todos    []Todo
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = services.Hash(u.Password)
	return
}

func (u *User) FindById(id int64) {
	mdb.First(&u, id)
}

func (u *User) FindByUsername(username string) {
	mdb.First(&u, "username = ?", username)
}

func (u *User) Create() {
	mdb.Create(&u)
}

func (u *User) Save() {
	db := db.Connect()
	db.Save(&u)
}

func (u *User) Destroy() {
	mdb.Delete(&u)
}

func (u *User) LoadTodos() {
	var todos []Todo
	mdb.Where("user_id = ?", u.Id).Find(&todos)
	u.Todos = todos
}
