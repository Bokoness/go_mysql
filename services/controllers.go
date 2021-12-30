package services

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ParseIdFromReq(r *http.Request) int64 {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	return id
}
