package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Username string `json:"username" schema:"username" `
	Email    string `gorm:"unique" schema:"email"  json:"email"`
	Password string `json:"password"`
	Roles    int64  `json:"roles"`
}
