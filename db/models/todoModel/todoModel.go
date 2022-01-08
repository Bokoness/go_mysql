package todoModel

import (
	"fmt"
	"go_mysql/db"
)

const model = "todos"

type Todo struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Uid     int64  `json:"uid"`
}

func (t Todo) Store () int64 {
	q := fmt.Sprintf(`INSERT INTO %s(title, content, uid) VALUES (?,?,?)`, model)
	var params []interface{}
	params = append(params, t.Title)
	params = append(params, t.Content)
	params = append(params, t.Uid)
	return db.Insert(q, params)
}

func (t Todo) Save() {
	q := fmt.Sprintf(`UPDATE %s SET title=?, content=? WHERE id=?`, model)
	var params []interface{}
	params = append(params, t.Title)
	params = append(params, t.Content)
	params = append(params, t.ID)
	db.Exec(q, params)
}

func (t Todo) Destroy(uid int64) {
	q := fmt.Sprintf("DELETE FROM %s WHERE id=? AND uid=?", model)
	var params []interface{}
	params = append(params, t.ID)
	params = append(params, uid)
	db.Exec(q, params)
}


func FindOneById(id int64) Todo {
	q := fmt.Sprintf("select * from %s where id=%d", model, id)
	rows := db.Find(q)
	var t Todo
	for rows.Next() {
		rows.Scan(&t.ID, &t.Title, &t.Content, &t.Uid)
	}
	return t
}

func FindManyById(uid int64) []Todo {
	var todos []Todo
	q := fmt.Sprintf("select * from %s where uid=%d", model, uid)
	rows := db.Find(q)
	for rows.Next() {
		var t Todo
		rows.Scan(&t.ID, &t.Title, &t.Content, &t.Uid)
		todos = append(todos, t)
	}
	return todos
}

func DeleteById(id int64, uid int64) {
	db.SafeDestroy(model, id, uid)
}
