package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/princetomar27/bookstore/pkg/config"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

// initialize the databse
func InitDB(){
	 
	fmt.Println("Connecting to the database...")
	config.ConnectToDatabase()
	db = config.GetDB()

	if db == nil {
		log.Fatal("Failed to initialize database")
	}
	
	fmt.Println("Migrating database schema...")
	db.AutoMigrate(&Book{})
	fmt.Println("Database initialized successfully.")

}	
func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByID(ID int64) (*Book, *gorm.DB){
	var getBook Book
	db := db.Where("ID=?", ID).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book{
	var book Book
	db.Where("ID=?",ID).Delete(book)
	return book
}