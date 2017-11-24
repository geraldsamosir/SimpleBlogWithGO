package model

import (
	"github.com/jinzhu/gorm"
)

//main model file for model users

// migarte function
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	return db
}
