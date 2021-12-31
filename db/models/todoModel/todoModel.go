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
