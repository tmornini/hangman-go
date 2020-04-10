package endpoints

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/interfaces"
)

type ErrorAnticipated struct{}

// RespondsTo implement interfaces.
func (ep ErrorAnticipated) RespondsTo(r *http.Request) bool {
	return r.URL.Path == "/error-anticipated"
}

// RespondsTo implement interfaces.Endpint
func (ep ErrorAnticipated) RespondTo(r *http.Request) (interfaces.Entity, error) {
	return errorAnticipatedEntity{}, nil
}

type errorAnticipatedEntity struct{}

// WriteResponseTo implement interfaces.Entity
func (ety errorAnticipatedEntity) WriteResponseTo(w http.ResponseWriter) {
	w.WriteHeader(503)
}
