package migration

import (
	table "github.com/Robinzon100/fiber/database/migrations/tables"
	"gorm.io/gorm"
)


var posts = []table.Post{
	{Title: "something", Description: "something even more"},
	{Title: "something else", Description: "something else even more"},
}


func Migrations(DB *gorm.DB) {
	DB.AutoMigrate(&table.Post{})
	// fmt.Println(DB)
	// for i := range posts {
	// 	DB.Create(&posts[i])
	// }
}
