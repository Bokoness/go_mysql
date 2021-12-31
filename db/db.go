package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func exec(q string, vals []interface{}) sql.Result {
	db, e := sql.Open("mysql", "root:321123@/go")
	ErrorCheck(e)
	stmt, e := db.Prepare(q)
	ErrorCheck(e)
	res, e := stmt.Exec(vals...)
	ErrorCheck(e)
	return res
}

func query(q string) *sql.Rows {
	db, e := sql.Open("mysql", "root:321123@/go")
	rows, e := db.Query(q)
	ErrorCheck(e)
	return rows
}

func Insert(m string, data map[string]string) int64 {
	res := exec(createInsertQuery(m, data))
	id, e := res.LastInsertId()
	ErrorCheck(e)
	return id
}

func FindById(m string, id int64) *sql.Rows {
	return query(fmt.Sprintf("select * from %s where id=%d", m, id))
}

func Find(m string, q string) *sql.Rows {
	return query(q)
}

func Destroy(m string, id int64) bool {
	res := exec(createDeleteQuery(m, id))
	a, e := res.RowsAffected()
	ErrorCheck(e)
	return a > 1
}

func SafeDestroy(m string, id int64, uid int64) {
	q := fmt.Sprintf("DELETE FROM %s WHERE id=? AND uid=?", m)
	var params []interface{}
	params = append(params, fmt.Sprint(id))
	params = append(params, fmt.Sprint(uid))
	exec(q, params)
}

func createInsertQuery(m string, data map[string]string) (string, []interface{}) {
	fields := "("
	qVals := "("
	var params []interface{}
	for k, v := range data {
		fields += k + ","
		qVals += "?,"
		params = append(params, v)
	}
	fields = fields[:len(fields)-1] + ")"
	qVals = qVals[:len(qVals)-1] + ")"
	q := fmt.Sprintf("INSERT INTO %s%s VALUES ", m, fields)
	q += qVals
	return q, params
}

func createDeleteQuery(m string, id int64) (string, []interface{}) {
	q := fmt.Sprintf("DELETE FROM %s WHERE id=?", m)
	var params []interface{}
	params = append(params, fmt.Sprint(id))
	return q, params
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func PingDB(db *sql.DB) {
	err := db.Ping()
	ErrorCheck(err)
}
