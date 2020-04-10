package endpoints

import (
	"errors"
	"net/http"

	"github.com/tmornini/udemy-hangman/entities"
	"github.com/tmornini/udemy-hangman/interfaces"
)

type ErrorUnanticipated struct{}

// RespondsToPathOf implement interfaces.Endpoint
func (ep ErrorUnanticipated) RespondsToPathOf(r *http.Request) bool {
	return r.URL.Path == "/error-unanticipated"
}

// RespondsTo implement interfaces.Endpoint
func (ep ErrorUnanticipated) RespondTo(r *http.Request) (interfaces.Entity, error) {
	if true {
		return nil, errors.New("this error is expected!")
	}

	return entities.GetErrorUnanticipatedEntity{}, nil
}

// Validate interfaces.Endpoint
func (ep ErrorUnanticipated) Validate(r *http.Request) (interfaces.Entity, error) {
	if r.Method != "GET" {
		return entities.MethodNotAllowed{}, nil
	}

	return nil, nil
}
