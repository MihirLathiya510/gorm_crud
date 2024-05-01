package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/MihirLathiya510/gorm_crud/pkg/mocks"
	"github.com/MihirLathiya510/gorm_crud/pkg/models"
	"github.com/google/uuid"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
    // Read to request body
    defer r.Body.Close()
    body, err := io.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
    }

    var book models.Book
    json.Unmarshal(body, &book)

    // Append to the Book mocks
    // book.Id = rand.Intn(100)
	book.Id = int(uuid.New().ID())
    mocks.Books = append(mocks.Books, book)

    // Send a 201 created response
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode("Created")
}
