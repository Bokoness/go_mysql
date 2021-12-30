package user

import (
	"go_mysql/db/models/userModel"
	"go_mysql/services"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Destroy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	u := userModel.User{ID: services.ParseIdFromReq(r)}
	u.Destroy()
}
