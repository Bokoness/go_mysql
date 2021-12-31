package auth

import (
	"encoding/json"
	"fmt"
	"go_mysql/db/models/userModel"
	"go_mysql/services"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var u userModel.User
	_ = json.NewDecoder(r.Body).Decode(&u)
	u.Save()
}

func Login(w http.ResponseWriter, r *http.Request) {
	var body userModel.User
	_ = json.NewDecoder(r.Body).Decode(&body)
	u := userModel.FindByUsername(body.UserName)
	fmt.Println(u.ID)

	if !services.ComparePasswords(u.Password, body.Password) {
		http.Error(w, "Forbidden", http.StatusUnauthorized)
		return
	}
	services.CreateCookieToken(u.ID, w)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	services.ClearAuth(w)
}
