package user

import (
	"go_mysql/db/models/userModel"
	"go_mysql/services"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Destroy(r *http.Request) {
	u := userModel.User{ID: services.ParseIdFromReq(r)}
	u.Destroy()
}


