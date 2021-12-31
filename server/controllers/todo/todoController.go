package todo

import (
	"encoding/json"
	"go_mysql/db/models/todoModel"
	"go_mysql/db/models/userModel"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Store(w http.ResponseWriter, r *http.Request) {
	var t todoModel.Todo
	_ = json.NewDecoder(r.Body).Decode(&t)
	u := userModel.GetActiveUser(r)
	t.Uid = u.ID
	t.Save()
}
