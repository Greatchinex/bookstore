package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/Greatchinex/bookstore/pkg/models"
	"github.com/Greatchinex/bookstore/pkg/utils"
)

type ResponseStatus struct {
	Status bool	`json:"status"`
	Message string `json:"message"`
}


func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := models.Book{}
	if err := utils.DecodeJson(r, &newBook); err != nil {
		fmt.Println(err)
	}

	book := newBook.CreateBook()
	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	allBooks := models.GetAllBooks()
	res, _ := json.Marshal(allBooks)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	Id := mux.Vars(r)["bookId"]
	ID, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Printf("Error converting string %v to integer\n", ID)
	}

	singleBook := models.GetBookById(ID)
	res, _ := json.Marshal(singleBook)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := models.Book{}
	if err := utils.DecodeJson(r, &updateBook); err != nil {
		fmt.Println(err)
	}

	Id := mux.Vars(r)["bookId"]
	ID, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Printf("Error converting string %v to integer\n", ID)
	}

	book := models.UpdateBook(ID, updateBook)
	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["bookId"]
	Id, err := strconv.ParseInt(params, 0, 0)
	if err != nil {
		fmt.Printf("Error converting string %v to integer\n", Id)
	}

	models.DeleteBook(Id)
	apiResponse := ResponseStatus{Status: true, Message: "Book deleted successfully"}
	res, _ := json.Marshal(apiResponse)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}