package table

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
