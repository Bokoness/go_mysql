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
	w.Header().Set("Content-Type", "application/json")
	var u userModel.User
	_ = json.NewDecoder(r.Body).Decode(&u)
	u.Save()

}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body userModel.User
	_ = json.NewDecoder(r.Body).Decode(&body)

	//TODO: fetch real user and compare passwords

	if !services.ComparePasswords(body.Password, "asdasd") {
		fmt.Println("NOT GOOD") //TODO: do something
	}

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"uid": u.ID})
	// hash, e := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	// if e != nil {
	// 	http.Error(w, "Cant log in", 500)
	// }
	// c := http.Cookie{Name: "uid", Value: hash, Path: "/"}
	// http.SetCookie(w, &c)
	// fmt.Fprintf(w, "User is now logged in, this is token: %s", hash)
}
