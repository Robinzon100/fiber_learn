package table

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID          uuid.UUID `json:"id",gorm:"type:uuid;primary_key;"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (post *Post) BeforeCreate(tx *gorm.DB) (err error) {
	post.ID = uuid.NewV4()
	return
}
