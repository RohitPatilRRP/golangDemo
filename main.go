package main

// @title Books API documentation
// @version 1.0.0
// @host localhost:5000
// @BasePath /api/books

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/RohitPatilRRP/golangDemo/docs"	
)

//Book struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Aurhor *Author `json:"author"`
}

//Author struct (Model)
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice book struct
var books []Book

// GetBooks ... Get all books
// @Summary Get all books
// @Description get all books
// @Tags Books
// @Success 200 {array} Book
// @Failure 404 {object} object
// @Router / [get]
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GET single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// CreateBook... Create Book
// @Summary Create new book based on paramters
// @Description Create new book
// @Tags Books
// @Accept json
// @Param user body Book true "Book Data"
// @Success 200 {object} object
// @Failure 400,500 {object} object
// @Router / [post]
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {

	r := mux.NewRouter()

	//Mock Data
	books = append(books, Book{ID: "1", Isbn: "123", Title: "Book One", Aurhor: &Author{Firstname: "Rohit", Lastname: "Patil"}})
	books = append(books, Book{ID: "2", Isbn: "132", Title: "Book Two", Aurhor: &Author{Firstname: "John", Lastname: "Miller"}})

	r.HandleFunc("/api/books/", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),//The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))

	log.Fatal(http.ListenAndServe(":8080", r))

}
