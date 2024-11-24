package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/princetomar27/bookstore/pkg/models"
	"github.com/princetomar27/bookstore/pkg/routes"
)

func main(){

	router := mux.NewRouter()
	models.InitDB();
	routes.RegisterBookStoreRoutes(router)

	http.Handle("/",router)
	fmt.Printf("Server starting at PORT : 9000")
	log.Fatal(http.ListenAndServe("localhost:9000", router))
	
}