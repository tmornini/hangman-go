package endpoints

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/interfaces"
)

type Root struct{}

// RespondsTo implement interfaces.Endpoint
func (ep Root) RespondsTo(r *http.Request) bool {
	return r.URL.Path == "/"
}

// RespondsTo implement interfaces.Endpoint
func (ep Root) RespondTo(r *http.Request) (interfaces.Entity, error) {
	return rootEntity{}, nil
}

type rootEntity struct{}

// WriteResponseTo implement interfaces.Entity
func (ety rootEntity) WriteResponseTo(w http.ResponseWriter) {
	w.WriteHeader(200)
}
