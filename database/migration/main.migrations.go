package migration

import (
	"github.com/Robinzon100/fiber/database/models"
	"gorm.io/gorm"
)




func Migrations(DB *gorm.DB) {
	DB.AutoMigrate(&models.Post{})
	// for i := range seeds.Posts {
	// 	DB.Create(&seeds.Posts[i])
	// }
}
