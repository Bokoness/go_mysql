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

func (t Todo) Save() {
	data := make(map[string]string)
	data["title"] = t.Title
	data["content"] = t.Content
	data["uid"] = fmt.Sprintf("%d", t.Uid)
	t.ID = db.Insert(model, data)
}

func (t Todo) Destroy() {
	db.Destroy(model, t.ID)
}

func FindOneById(id int64) Todo {
	q := fmt.Sprintf("select * from %s where id=%d", model, id)
	rows := db.Find(model, q)
	var t Todo
	for rows.Next() {
		rows.Scan(&t.ID, &t.Title, &t.Content, &t.Uid)
	}
	return t
}

func FindManyById(uid int64) []Todo {
	var todos []Todo
	q := fmt.Sprintf("select * from %s where uid=%d", model, uid)
	rows := db.Find(model, q)
	for rows.Next() {
		var t Todo
		rows.Scan(&t.ID, &t.Title, &t.Content, &t.Uid)
		todos = append(todos, t)
	}
	return todos
}

func DeleteById(id int64, uid int64) {
	// db.Destroy(model, id)

	db.SafeDestroy(model, id, uid)
}
