package models

import "go_mysql/db"

type Model interface {
	FindById(int64)
	Create()
	Save()
}

var mdb = db.Connect()
