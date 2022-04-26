package services

import (
	"encoding/json"
	"net/http"
	"strconv"

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

func BodyIntoMap(r *http.Request) (map[string]string, error) {
	m := make(map[string]string)
	e := json.NewDecoder(r.Body).Decode(&m)
	if e != nil {
		return nil, e
	}
	return m, nil
}

func WriteDataIntoResponse(d interface{}, w http.ResponseWriter) {
	j, e := json.Marshal(d)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(j)
}

func ErrorCheck(err error, status int, w http.ResponseWriter) {
	if err != nil {
		w.WriteHeader(400)
	}
}
