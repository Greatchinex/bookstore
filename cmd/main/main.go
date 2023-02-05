package main

import (
	"log"
	"net/http"

	"github.com/Greatchinex/bookstore/pkg/routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)


func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:9010", router))
}