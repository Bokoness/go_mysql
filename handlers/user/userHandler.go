package user

import (
	"go_mysql/db/models"
	"go_mysql/services"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Destroy(r *http.Request) {
	var u models.User
	u.FindById(services.ParseIdFromReq(r))
	u.Destroy()
}
