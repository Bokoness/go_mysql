package todo

import (
	"encoding/json"
	"fmt"
	"go_mysql/db/models"
	"go_mysql/server/middleware"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func Store(w http.ResponseWriter, r *http.Request) {
	//TODO: fetch User from c
	c := r.Context().Value("user")
	fmt.Println(c)
	var t models.Todo
	_ = json.NewDecoder(r.Body).Decode(&t)
	u, _ := middleware.FetchUserFromCookie(r)
	t.UserID = u.Id
	t.Create()
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	u, _ := middleware.FetchUserFromCookie(r)
	todos := u.LoadTodos()
	j, e := json.Marshal(todos)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(j)
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	var todo models.Todo
	todo.FindById(id)
	j, e := json.Marshal(todo)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(j)
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var t models.Todo
	if e := json.NewDecoder(r.Body).Decode(&t); e != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	u, _ := middleware.FetchUserFromCookie(r)
	t.Id = id
	t.UserID = u.Id
	t.Save()
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	u, _ := middleware.FetchUserFromCookie(r)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	var t models.Todo
	t.FindById(id)
	if t.Id != u.Id {
		//TODO: throw error
		return
	}
	t.Destroy()
}
