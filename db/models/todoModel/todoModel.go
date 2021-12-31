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

func FindManyById(uid int64) []Todo {
	var todos []Todo
	q := fmt.Sprintf("select * from %s where uid=\"%d\"", model, uid)
	rows := db.Find(model, q)
	for rows.Next() {
		var t Todo
		rows.Scan(&t.ID, &t.Title, &t.Content, &t.Uid)
		todos = append(todos, t)
	}
	return todos
}
