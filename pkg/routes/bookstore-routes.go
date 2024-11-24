package routes

import (
	"github.com/gorilla/mux"
	"github.com/princetomar27/bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func (router *mux.Router){
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

}

