package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/uday778/LibraryManagementSystem-with-gorillaMux/middleware"
	"github.com/uday778/LibraryManagementSystem-with-gorillaMux/routes"
)

// login
// signUp
// order
// see the list of books he ordered
// see what books are available
// owner should able to add new books
// owner should able to delete books


func main()  {
	router:= mux.NewRouter()
	routes.BookRoutes(router)
	routes.UserRoutes(router)
	routes.OrderRoutes(router)
	router.Use(middleware.TrackNumberOfRequests)

	http.ListenAndServe(":4000", router)
}



