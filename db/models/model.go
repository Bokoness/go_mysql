package models

import (
	"go_mysql/db"
	"io"
)

type Model interface {
	Index(int64)
	Show(int64)
	Create(int64)
	Save()
	Destroy()
	Decode(io.Reader)
	EncodeOne() ([]byte, error)
	EncodeMany() ([]byte, error)
}

var mdb = db.Connect()
