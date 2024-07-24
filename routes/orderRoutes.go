package routes

import (
	"github.com/gorilla/mux"

	"github.com/uday778/LibraryManagementSystem-with-gorillaMux/controllers"
)

func OrderRoutes(router *mux.Router) {
	router.HandleFunc("/orderBook/{bookId}",controllers.OrderBook)

	router.HandleFunc("/listallorderedbooks",controllers.Listallorderedbooks)
}