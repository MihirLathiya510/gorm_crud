package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/MihirLathiya510/gorm_crud/pkg/mocks"
	"github.com/MihirLathiya510/gorm_crud/pkg/models"
	"github.com/gorilla/mux"
)

func UpdateBook(w http.ResponseWriter, r *http.Request) {
    // Read dynamic id parameter
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    // Read request body
    defer r.Body.Close()
    body, err := io.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
    }

    var updatedBook models.Book
    json.Unmarshal(body, &updatedBook)

    // Iterate over all the mock Books
   	// Iterate over all the mock Books
	for index, book := range mocks.Books {
		if book.Id == id {
			// Update the specified field
			if updatedBook.Title != "" {
				book.Title = updatedBook.Title
			}
			if updatedBook.Author != "" {
				book.Author = updatedBook.Author
			}
			if updatedBook.Desc != "" {
				book.Desc = updatedBook.Desc
			}

			// Update the book in the slice
			mocks.Books[index] = book

			// Send a response
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Updated")
			return
		}
	}

}