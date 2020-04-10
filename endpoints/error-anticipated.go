package endpoints

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/entities"
	"github.com/tmornini/udemy-hangman/interfaces"
)

type ErrorAnticipated struct{}

// RespondsToPathOf implement interfaces.Endpoint
func (ep ErrorAnticipated) RespondsToPathOf(r *http.Request) bool {
	return r.URL.Path == "/error-anticipated"
}

// RespondTo implement interfaces.Endpoint
func (ep ErrorAnticipated) RespondTo(r *http.Request) (interfaces.Entity, error) {
	return entities.GetErrorAnticipated{}, nil
}

// Validate interfaces.Endpoint
func (ep ErrorAnticipated) Validate(r *http.Request) (interfaces.Entity, error) {
	if r.Method != "GET" {
		return entities.MethodNotAllowed{}, nil
	}

	return nil, nil
}
