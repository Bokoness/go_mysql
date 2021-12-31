package todo

import (
	"encoding/json"
	"go_mysql/db/models/todoModel"
	"go_mysql/server/middleware"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Store(r *http.Request) {
	var t todoModel.Todo
	_ = json.NewDecoder(r.Body).Decode(&t)
	u, _ := middleware.FetchUserFromCookie(r)
	t.Uid = u.ID
	t.Save()
}

func Index(r *http.Request, w http.ResponseWriter) []byte {
	w.Header().Set("Content-Type", "application/json")
	u, _ := middleware.FetchUserFromCookie(r)
	todos := todoModel.FindManyById(u.ID)
	todoJson, e := json.Marshal(todos)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	return todoJson
}
