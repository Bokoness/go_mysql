package auth

import (
	"encoding/json"
	"go_mysql/db/models"
	"go_mysql/services"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var u models.User
	_ = json.NewDecoder(r.Body).Decode(&u)
	u.Create()
}

func Login(w http.ResponseWriter, r *http.Request) {
	var body models.User
	_ = json.NewDecoder(r.Body).Decode(&body)
	var u models.User
	u.FindByUsername(body.Username)
	if !services.ComparePasswords(u.Password, body.Password) {
		http.Error(w, "Forbidden", http.StatusUnauthorized)
		return
	}
	services.CreateCookieToken(u.Id, w)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	services.ClearAuth(w)
}
