package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "root:321123@tcp(127.0.0.1:3306)/go_mysql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("No connection to database")
	}
	return db
}

func Exec(q string, params []interface{}) sql.Result {
	db, e := sql.Open("mysql", "root:321123@/go_mysql")
	defer db.Close()
	ErrorCheck(e)
	stmt, e := db.Prepare(q)
	ErrorCheck(e)
	res, e := stmt.Exec(params...)
	ErrorCheck(e)
	return res
}

func Query(q string) *sql.Rows {
	db, e := sql.Open("mysql", "root:321123@/go_mysql")
	defer db.Close()
	rows, e := db.Query(q)
	ErrorCheck(e)
	return rows
}

func Insert(q string, params []interface{}) int64 {
	res := Exec(q, params)
	id, e := res.LastInsertId()
	ErrorCheck(e)
	return id
}

func FindById(m string, id int64) *sql.Rows {
	return Query(fmt.Sprintf("select * from %s where id=%d", m, id))
}

func Find(q string) *sql.Rows {
	return Query(q)
}

func SafeDestroy(m string, id int64, uid int64) {
	q := fmt.Sprintf("DELETE FROM %s WHERE id=? AND uid=?", m)
	var params []interface{}
	params = append(params, fmt.Sprint(id))
	params = append(params, fmt.Sprint(uid))
	Exec(q, params)
}

//Private functions

//TODO: upgrade

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}
