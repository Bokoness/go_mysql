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
