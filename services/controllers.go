package services

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func ParseIdFromReq(r *http.Request) int64 {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	return id
}

func CreateCookieToken(uid int64, w http.ResponseWriter) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"uid": uid})
	hash, e := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if e != nil {
		http.Error(w, "Cant log in", 500)
	}
	c := http.Cookie{Name: "uid", Value: hash, Path: "/"}
	http.SetCookie(w, &c)
	fmt.Fprintf(w, "User is now logged in, this is token: %s", hash)
}

func ClearAuth(w http.ResponseWriter) {
	c := http.Cookie{
		Name:   "uid",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, &c)
}
