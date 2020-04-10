package endpoints

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/interfaces"
	"github.com/tmornini/udemy-hangman/responses"
	"github.com/tmornini/udemy-hangman/uuid"
)

// Endpoint

type PostGames struct{}

// RespondsToPathOf implement interfaces.Endpoint
func (ep PostGames) RespondsToPathOf(r *http.Request) bool {
	return r.URL.Path == "/games/"
}

// RespondsToMethodOf implement interfaces.Endpoint
func (ep PostGames) RespondsToMethodOf(r *http.Request) bool {
	return r.Method == "POST"
}

// Process implement interfaces.Endpoint
func (ep PostGames) Process(r *http.Request) (interfaces.Responsible, error) {
	return responses.PostGames{ID: uuid.Generate()}, nil
}
