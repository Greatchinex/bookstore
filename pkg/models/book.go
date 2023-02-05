package models

import (
	"github.com/Greatchinex/bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
type Book struct {
	gorm.Model // Auto gives you db columns id, and timestamps
	Name string `gorm:"name" json:"name"`
	Author string `gorm:"author" json:"author"`
	Publication string `gorm:"publication" json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}