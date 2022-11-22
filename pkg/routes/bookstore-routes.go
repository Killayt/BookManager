package routes

import (
	"github.com/Killayt/BookManager/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook.Methods("POST"))
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{id:[0-9]+}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id:[0-9]+}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id:[0-9]+}", controllers.DeleteBook).Methods("DELETE")

}
