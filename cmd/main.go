package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MihirLathiya510/gorm_crud/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Hello World
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode("Hello World")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// CRUD operations
	router.HandleFunc("/books", handlers.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/book/{id}", handlers.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books", handlers.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/book/{id}", handlers.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/book/{id}", handlers.DeleteBook).Methods(http.MethodDelete)
	
	// Start the server
	log.Println("API is running!")
	err := http.ListenAndServe(":4000", router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
