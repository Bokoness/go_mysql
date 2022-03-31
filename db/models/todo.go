package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  int64  `json:"userId"`
	User    User
}

func (t *Todo) FindById(id int64) {
	mdb.First(&t, id)
}

func (t *Todo) Create() {
	mdb.Create(&t)
}

func (t *Todo) Save() {
	mdb.Save(&t)
}

func (t *Todo) Destroy() {
	mdb.Delete(&t)
}

func (t Todo) Index() []Todo {
	var todos []Todo
	mdb.Find(&todos)
	return todos
}
