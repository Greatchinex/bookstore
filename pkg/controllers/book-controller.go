package controllers

import (
	"fmt"
	"net/http"
)


func CreateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Book Method")
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Multiple Books Method")
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Single Book Method")
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Existing Book Method")
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Single Book Method")
}