package endpoints

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/interfaces"
	"github.com/tmornini/udemy-hangman/responses"
)

// Endpoint

type GetRoot struct{}

// RespondsToPathOf implement interfaces.Endpoint
func (ep GetRoot) RespondsToPathOf(r *http.Request) bool {
	return r.URL.Path == "/"
}

// RespondsToMethodOf implement interfaces.Endpoint
func (ep GetRoot) RespondsToMethodOf(r *http.Request) bool {
	return r.Method == "GET"
}

// Process implement interfaces.Endpoint
func (ep GetRoot) Process(r *http.Request) (interfaces.Responsible, error) {
	return responses.GetRoot{}, nil
}
