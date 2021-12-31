package userModel

import (
	"fmt"
	"go_mysql/db"
	Services "go_mysql/services"
)

const model = "users"

type User struct {
	ID       int64  `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (u User) Save() {
	u.Password = Services.Hash(u.Password)
	data := make(map[string]string)
	data["username"] = u.UserName
	data["password"] = u.Password
	u.ID = db.Insert(model, data)
}

func (u User) Destroy() {
	db.Destroy(model, u.ID)
}

func FindByUsername(username string) User {
	var u User
	q := fmt.Sprintf("select * from %s where username=\"%s\"", model, username)
	rows := db.Find(model, q)
	for rows.Next() {
		rows.Scan(&u.ID, &u.UserName, &u.Password)
	}
	return u
}

func FindById(uid int64) User {
	var u User
	rows := db.FindById(model, uid)
	for rows.Next() {
		rows.Scan(&u.ID, &u.UserName, &u.Password)
	}
	return u
}
