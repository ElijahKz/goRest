package main

import(
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"fmt"
)

type Book struct{
	Id int `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Descripcion string `json:"descripcion,omitempty"`
	Author string `json:"author,omitempty"`
}

var books []Book
//curl -i http://localhost:3000/books
func GetBooksEndpoint(w http.ResponseWriter, req *http.Request) {
json.NewEncoder(w).Encode(books)
}
//curl -i http://localhost:3000/books/1

func GetBookEndpoint(w http.ResponseWriter, req *http.Request) {	
	items := mux.Vars(req)	
	for _, item := range books{
		//Convirtiendo en entero el valor del id entrante, el cual viene como cadena
		idItem, err := strconv.Atoi(items["id"])
		if err != nil {
			fmt.Printf("%+v ", err)
		}
		if item.Id == idItem{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}
//curl -i -H "Content-Type: application/json" -X POST -d '{"title":"aamateur.docx"}' http://localhost:3000/books
func CreateBookEndpoint(w http.ResponseWriter, req *http.Request) {
	var book Book 
	_= json.NewDecoder(req.Body).Decode(&book)
	if book.Title != " "{
		//Se verfica que el titulo no está vacío y se agrega un nuevo libro autoincrementando
		books = append(books,  Book{Id:books[len(books)-1].Id+1, Title:book.Title, Descripcion: book.Descripcion , Author:book.Author})
		json.NewEncoder(w).Encode(books)
	}
}
// curl -i -H "Content-Type: application/json" -X DELETE http://localhost:3000/books/2
func DeleteBookEndpoint(w http.ResponseWriter, req *http.Request) {
	items := mux.Vars(req)
	for index, item := range books{
		idItem, err := strconv.Atoi(items["id"])
		if err != nil {
			fmt.Printf("%+v ", err)
		}
		if item.Id == idItem{
			books = append(books[:index], books[index + 1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
// curl -i -H "Content-Type: application/json" -X PUT -d '{"author":"amael"}' http://localhost:3000/books/2
func UpdateBookEndpoint (w http.ResponseWriter, req *http.Request) {
	items := mux.Vars(req)
	// Creando un objeto book a partir del request
	var book Book 
	_= json.NewDecoder(req.Body).Decode(&book)
	for indice, item := range books{
		//Convirtiendo en entero el valor del id entrante, el cual viene como cadena
		
		idItem, err := strconv.Atoi(items["id"])
		if err != nil {
			fmt.Printf("%+v ", err)
		}
		if item.Id == idItem{
			books[indice].Id = item.Id 
			if( book.Title != ""){ books[indice].Title = book.Title}		
			
			if( book.Descripcion != "" ){ books[indice].Descripcion = book.Descripcion}

			if( book.Author != "" ) { books[indice].Author = book.Author}
			
			
			json.NewEncoder(w).Encode(books)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func main() {
	router :=mux.NewRouter()
	// Creando en memoria los datos
	books = append(books, Book{Id:1, Title:"cien años de soledad", Descripcion: "90's ganador . ." , Author:"gabo"})
	//Manejando las rutas con la biblioteca http mux
	router.HandleFunc("/books", GetBooksEndpoint).Methods("GET")
	router.HandleFunc("/books/{id}", GetBookEndpoint).Methods("GET")
	router.HandleFunc("/books", CreateBookEndpoint).Methods("POST")
	router.HandleFunc("/books/{id}", DeleteBookEndpoint).Methods("DELETE")
	router.HandleFunc("/books/{id}", UpdateBookEndpoint).Methods("PUT")
	// corriendo servidor http
	log.Fatal(http.ListenAndServe(":3000", router))
}
