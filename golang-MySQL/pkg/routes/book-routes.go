package routes

import (
	"github.com/Go_Projects/golang-MySQL/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/getBooks", controllers.GetBook).Methods("GET")
	router.HandleFunc("/getBook/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/updateBook/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/deleteBook/{id}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/createBook", controllers.CreateBook).Methods("POST")
}
