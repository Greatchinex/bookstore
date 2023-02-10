package models

import (
	"github.com/Greatchinex/bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
type Book struct {
	gorm.Model // Auto gives you db columns id, and timestamps along with deleted_at
	Name string `gorm:"name" json:"name"`
	Author string `gorm:"author" json:"author"`
	Publication string `gorm:"publication" json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) Book {
	var singleBook Book
	db.Where("ID=?", Id).Find(&singleBook)
	return singleBook
}

func DeleteBook(Id int64) bool {
	var singleBook Book
	db.Where("ID=?", Id).Delete(&singleBook)
	return true
} 