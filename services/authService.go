package services

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateCookieToken(uid int64, w http.ResponseWriter) {
	expireDays := 3
	d, _ := time.ParseDuration(fmt.Sprintf("%dh", expireDays*24))
	claims := customClaims{
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(d).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	hash, e := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if e != nil {
		http.Error(w, "Cant log in", 500)
	}
	c := http.Cookie{Name: "uid", Value: hash, Path: "/"}
	http.SetCookie(w, &c)
}

func CreateToken() {

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
