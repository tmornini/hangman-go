package endpoints

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/entities"
	"github.com/tmornini/udemy-hangman/interfaces"
)

// Endpoint

type ErrorAnticipated struct{}

// RespondsToPathOf implement interfaces.Endpoint
func (ep ErrorAnticipated) RespondsToPathOf(r *http.Request) bool {
	return r.URL.Path == "/error-anticipated"
}

// RespondTo implement interfaces.Endpoint
func (ep ErrorAnticipated) RespondTo(r *http.Request) (interfaces.Entity, error) {
	return errorAnticipatedEntity{}, nil
}

// Validate interfaces.Endpoint
func (ep ErrorAnticipated) Validate(r *http.Request) (interfaces.Entity, error) {
	if r.Method != "GET" {
		return entities.MethodNotAllowed{}, nil
	}

	return nil, nil
}

// Entity

type errorAnticipatedEntity struct{}

// WriteResponseTo implement interfaces.Entity
func (ety errorAnticipatedEntity) WriteResponseTo(w http.ResponseWriter) {
	w.WriteHeader(503)
}
