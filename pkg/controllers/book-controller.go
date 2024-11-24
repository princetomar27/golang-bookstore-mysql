package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/princetomar27/bookstore/pkg/models"
	"github.com/princetomar27/bookstore/pkg/utils"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter,r *http.Request){

	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks) // convert it to JSON

	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)

	// To send data to frontend / POSTMAN
	w.Write(res)

}

func GetBookByID(w http.ResponseWriter, r *http.Request){

	params := mux.Vars(r)
	bookId := params["bookId"]
	ID,err := strconv.ParseInt(bookId, 0,0) // parse to INT64
	if err != nil{
		fmt.Print("Error while Parsing")
	}
	bookDetails, _ := models.GetBookByID(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateBook(w http.ResponseWriter, r *http.Request){

	newBook := &models.Book{}
	// user has a request body, parse it to db from JSON body
	utils.ParseBody(r, newBook)
	b := newBook.CreateBook()
	res,_ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){

	params := mux.Vars(r)
	bookID := params["bookId"]
	// parse bookId to int64
	ID, err := strconv.ParseInt(bookID,0,0)
	if err != nil{
		fmt.Println("Error while Parsing book id")
	}
	bookDetails := models.DeleteBook(ID)
	res,_ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request){

	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	params := mux.Vars(r)
	bookId := params["bookId"]

	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("Error while parsing")
	}

	bookDetails, db := models.GetBookByID(ID)
	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}
	

	// save in the db
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	 
}