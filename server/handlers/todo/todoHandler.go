package todo

import (
	"encoding/json"
	"go_mysql/db/models/todoModel"
	"go_mysql/server/middleware"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
)

func Store(r *http.Request) {
	var t todoModel.Todo
	_ = json.NewDecoder(r.Body).Decode(&t)
	u, _ := middleware.FetchUserFromCookie(r)
	t.Uid = u.ID
	t.Store()
}

func Index(r *http.Request, w http.ResponseWriter) []byte {
	w.Header().Set("Content-Type", "application/json")
	u, _ := middleware.FetchUserFromCookie(r)
	todos := todoModel.FindManyById(u.ID)
	j, e := json.Marshal(todos)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	return j
}

func Show(w http.ResponseWriter, p martini.Params) []byte {
	w.Header().Set("Content-Type", "application/json")
	id, e := strconv.ParseInt(p["id"], 10, 64)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	todo := todoModel.FindOneById(id)
	j, e := json.Marshal(todo)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	return j
}

func Update(r *http.Request, w http.ResponseWriter, p martini.Params) {
	var t todoModel.Todo
	if e := json.NewDecoder(r.Body).Decode(&t) ; e != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	id, e := strconv.ParseInt(p["id"], 10, 64)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	u, _ := middleware.FetchUserFromCookie(r)
	t.ID = id
	t.Uid = u.ID
	t.Save()
}

func Destroy(r *http.Request, w http.ResponseWriter, p martini.Params) {
	id, e := strconv.ParseInt(p["id"], 10, 64)
	u, _ := middleware.FetchUserFromCookie(r)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	todoModel.DeleteById(id, u.ID)
}
