package seeders

import (
	"fmt"
	"go_mysql/db/models"

	"gorm.io/gorm"
)

func TodoSeed(db *gorm.DB) {
	todos := []models.Todo{}
	var j int64 = 1
	for ; j < 30; j++ {
		tname := fmt.Sprintf("todo%d", j)
		todos = append(todos, models.Todo{
			Title:   tname,
			Content: "content",
			UserID:  j / 3,
		})
	}
	db.Create(&todos)
}
