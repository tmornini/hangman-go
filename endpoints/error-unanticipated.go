package endpoints

import (
	"errors"
	"net/http"

	"github.com/tmornini/udemy-hangman/interfaces"
)

type ErrorUnanticipated struct{}

// RespondsTo implement interfaces.
func (ep ErrorUnanticipated) RespondsTo(r *http.Request) bool {
	return r.URL.Path == "/error-unanticipated"
}

// RespondsTo implement interfaces.Endpoint
func (ep ErrorUnanticipated) RespondTo(r *http.Request) (interfaces.Entity, error) {
	if true {
		return nil, errors.New("this error is expected!")
	}

	return errorUnanticipatedEntity{}, nil
}

type errorUnanticipatedEntity struct{}

// WriteResponseTo implement interfaces.Entity
func (ety errorUnanticipatedEntity) WriteResponseTo(w http.ResponseWriter) {
	w.WriteHeader(200)
}
