package models

import (
	"encoding/json"
	"io"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  int64  `json:"userId"`
	User    User
}

type TodoModule struct {
	Todo  Todo
	Todos []Todo
}

func (t *TodoModule) Show(id int64) {
	mdb.First(&t.Todo, id)
}

func (t *TodoModule) Create(uid int64) {
	t.Todo.UserID = uid
	mdb.Create(&t.Todo)
}

func (t *TodoModule) Save() {
	mdb.Save(&t.Todo)
}

func (t *TodoModule) Destroy() {
	mdb.Delete(&t.Todo)
}

func (t *TodoModule) Index(uid int64) {
	mdb.Find(&t.Todos)
}

func (t *TodoModule) Decode(body io.Reader) {
	_ = json.NewDecoder(body).Decode(&t)
}

func (t *TodoModule) EncodeOne() ([]byte, error) {
	j, err := json.Marshal(t.Todo)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (t *TodoModule) EncodeMany() ([]byte, error) {
	j, err := json.Marshal(t.Todos)
	if err != nil {
		return nil, err
	}
	return j, nil
}
