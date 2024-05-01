package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MihirLathiya510/gorm_crud/pkg/mocks"
	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the book ID from the request
	vars := mux.Vars(r)
	id := vars["id"]

	// Convert the ID from string to int
	bookID, err := strconv.Atoi(id)
	if err != nil {
		// If the ID is not a valid integer, return bad request
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid book ID")
		return
	}

	// Find the book from the mocks
	for _, book := range mocks.Books {
		if book.Id == bookID {
			// Book found, return it
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	// Book not found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Book not found")
}
