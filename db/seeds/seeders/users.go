package seeders

import (
	"fmt"
	"go_mysql/db/models"

	"gorm.io/gorm"
)

func UserSeed(db *gorm.DB) {

	users := []models.User{
		{
			Username: "bokoness",
			Password: "321123",
		},
	}
	for i := 0; i < 10; i++ {
		uname := fmt.Sprintf("name%d", i)
		users = append(users, models.User{
			Username: uname,
			Password: "321123",
		})
	}
	db.Create(&users)
}
