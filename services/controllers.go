package services

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type customClaims struct {
	Uid int64 `json:"uid"`
	jwt.StandardClaims
}

func ParseIdFromReq(r *http.Request) int64 {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	return id
}

func CreateCookieToken(uid int64, w http.ResponseWriter) {
	expireDays := 3
	d, _ := time.ParseDuration(fmt.Sprintf("%dh", expireDays*24))
	claims := customClaims{
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(d).Unix()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	hash, e := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if e != nil {
		http.Error(w, "Cant log in", 500)
	}
	c := http.Cookie{Name: "uid", Value: hash, Path: "/"}
	http.SetCookie(w, &c)
	fmt.Fprint(w, "User is now logged in")
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
