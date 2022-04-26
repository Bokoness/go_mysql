package handlers

import (
	"encoding/json"
	"go_mysql/db/models"
	"go_mysql/middleware"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	Model models.Model
}

func (h Handler) Store(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(h.Model)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	u, _ := middleware.FetchUserFromCookie(r)
	h.Model.Create(u.Id)
}

func (h Handler) Index(w http.ResponseWriter, r *http.Request) {
	u := middleware.FetchUserFromCtx(r)
	h.Model.Index(u.Id)
	j, err := h.Model.EncodeMany()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(j)
}

func (h Handler) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	h.Model.Show(id)
	j, e := json.Marshal(h.Model)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(j)
}

//TODO: adjust this handler to update all kind of models
func (h Handler) Update(w http.ResponseWriter, r *http.Request) {
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
	// t.Save()
}

func (h Handler) Destroy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, e := strconv.ParseInt(vars["id"], 10, 64)
	// u, _ := middleware.FetchUserFromCookie(r)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	h.Model.Show(id)
	h.Model.Destroy()
}
