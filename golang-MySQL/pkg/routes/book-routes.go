package routes

import (
	"github.com/Go_Projects/golang-MySQL/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/getBooks", controllers.GetBook).Methods("Get")
	router.HandleFunc("/getBook/{id}", controllers.GetBookById).Methods("Get")
	router.HandleFunc("/updateBook/{id}", controllers.UpdateBook).Methods("Put")
	router.HandleFunc("/deleteBook/{id}", controllers.DeleteBook).Methods("Delete")
	router.HandleFunc("/createBook", controllers.DeleteBook).Methods("Post")
}
