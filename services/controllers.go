package services

import (
	"encoding/json"
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

func ClearAuth(w http.ResponseWriter) {
	c := http.Cookie{
		Name:   "uid",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, &c)
}

func BodyIntoMap(r *http.Request) (map[string]string, error) {
	m := make(map[string]string)
	e := json.NewDecoder(r.Body).Decode(&m)
	if e != nil {
		return nil, e
	}
	return m, nil
}
