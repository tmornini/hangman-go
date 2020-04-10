package endpoints

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/entities"
	"github.com/tmornini/udemy-hangman/interfaces"
)

// Endpoint

type Root struct{}

// RespondsToPathOf implement interfaces.Endpoint
func (ep Root) RespondsToPathOf(r *http.Request) bool {
	return r.URL.Path == "/"
}

// RespondTo implement interfaces.Endpoint
func (ep Root) RespondTo(r *http.Request) (interfaces.Entity, error) {
	return rootEntity{}, nil
}

// Validate interfaces.Endpoint
func (ep Root) Validate(r *http.Request) (interfaces.Entity, error) {
	if r.Method != "GET" {
		return entities.MethodNotAllowed{}, nil
	}

	return nil, nil
}

// Entity

type rootEntity struct{}

// WriteResponseTo implement interfaces.Entity
func (ety rootEntity) WriteResponseTo(w http.ResponseWriter) error {
	w.WriteHeader(200)

	return nil
}
