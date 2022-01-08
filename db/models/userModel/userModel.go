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

func (u User) Store () int64 {
	q := fmt.Sprintf(`INSERT INTO %s(username, password) VALUES (?,?)`, model)
	var params []interface{}
	params = append(params, u.UserName)
	params = append(params, u.Password)
	return db.Insert(q, params)
}

func (u User) Save() {
	q := fmt.Sprintf(`UPDATE %s SET username=?, password=? WHERE id=?`, model)
	var params []interface{}
	u.Password = Services.Hash(u.Password)
	params = append(params, u.UserName)
	params = append(params, u.Password)
	params = append(params, u.ID)
	db.Exec(q, params)
}

func (u User) Destroy() {
	q := fmt.Sprintf("DELETE FROM %s WHERE id=?", model)
	var params []interface{}
	params = append(params, u.ID)
	db.Exec(q, params)
}

func FindByUsername(username string) User {
	var u User
	q := fmt.Sprintf("select * from %s where username=\"%s\"", model, username)
	rows := db.Find(q)
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
