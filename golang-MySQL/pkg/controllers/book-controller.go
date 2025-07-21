package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Go_Projects/golang-MySQL/pkg/models"
	"github.com/Go_Projects/golang-MySQL/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	utils.JsonResponse(w, 200, "books fetched successfully", newBooks)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
		utils.JsonResponse(w, 400, err.Error(), nil)
	}
	bookDetails, _ := models.GetBookById(ID)
	utils.JsonResponse(w, 200, "book fetched successfully", bookDetails)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	err := json.NewDecoder(r.Body).Decode(&createBook)
	if err != nil {
		utils.JsonResponse(w, http.StatusBadRequest, "Invalid data", nil)
		return
	}
	b := createBook.CreateBook()
	utils.JsonResponse(w, 201, "book creted", b)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
		utils.JsonResponse(w, 400, err.Error(), nil)
	}
	b := models.DeleteBookById(ID)
	utils.JsonResponse(w, 200, "book deleted successfully", b)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
		utils.JsonResponse(w, 400, err.Error(), nil)
	}
	var updateBook = &models.Book{}
	err = json.NewDecoder(r.Body).Decode(&updateBook)
	if err != nil {
		utils.JsonResponse(w, 400, err.Error(), nil)
		return
	}
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	utils.JsonResponse(w, 200, "updated successfully", bookDetails)
}
