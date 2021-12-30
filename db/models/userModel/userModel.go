package userModel

import (
	"go_mysql/db"
	"go_mysql/services"
)

const model = "users"

type User struct {
	ID       int64  `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (u User) Save() {
	u.Password = services.Hash(u.Password)
	data := make(map[string]string)
	data["username"] = u.UserName
	data["password"] = u.Password
	u.ID = db.Insert(model, data)
}

func (u User) Destroy() {
	db.Destroy(model, u.ID)
}