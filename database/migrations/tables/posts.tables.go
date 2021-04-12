package table

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
