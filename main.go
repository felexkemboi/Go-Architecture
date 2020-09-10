package main

import (
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"encoding/json"
	"github.com/gorilla/mux"
)


//Book structs (Model)
type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
    Title string `json:"title"`
	Author *Author `json:"author"`
}


//Author struct
type Author struct {
	Firstname string `json:"Firstame"`
	Lastname string `json:"Lastname"`
}


//Init books variable as a slice Book struct
var books []Book



//Get all books
func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)
}

//Get single book
func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r) //get params

	// Loop through the books
	for _, item := range books {
		if item.ID == params["id"]{
		json.NewEncoder(w).Encode(item)
		return 
		}
	}
	json.NewEncoder(w).Encode(&Book{})

}

//create new book
func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) //Mock ID
	books = append(books,book)


}

//Update book
func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index ,item := range books{
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
}

//Delete book
func updateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index ,item := range books{
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = strconv.Itoa(rand.Intn(10000000)) //Mock ID
			books = append(books,book)
			break
		}
	}
}



func main(){
	r := mux.NewRouter()

	//Mock data to 
	books = append(books, Book{ID :"1", Isbn: "4487",Title: "Book one", Author : &Author { Firstname : "John", Lastname: "Kemboi"}})
	books = append(books, Book{ ID :"2", Isbn: "4337", Title: "Book Two", Author : &Author{ Firstname : "Ian",  Lastname: "Githu"}})

	//Route Handlers to establish endpoints

	r.HandleFunc("/api/books",            getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}",       getBook).Methods("GET")
	r.HandleFunc("/api/books",            createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}",       getBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}",       deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000",r))

}