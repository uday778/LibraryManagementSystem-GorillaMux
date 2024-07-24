package routes

import (
	"github.com/gorilla/mux"

	"github.com/uday778/LibraryManagementSystem-with-gorillaMux/controllers"
)

func BookRoutes(router *mux.Router) {
	router.HandleFunc("/getBooks",controllers.GetBooks)

	router.HandleFunc("/getBook/{bookId}",controllers.GetBook)
	
	router.HandleFunc("/deletebook/{bookId}",controllers.DeleteBook)

	router.HandleFunc("/addBook",controllers.AddBook)
}
