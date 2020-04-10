package endpoints

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/entities"
	"github.com/tmornini/udemy-hangman/interfaces"
	"github.com/tmornini/udemy-hangman/uuid"
)

// Endpoint

type Games struct{}

// RespondsToPathOf implement interfaces.Endpoint
func (ep Games) RespondsToPathOf(r *http.Request) bool {
	return r.URL.Path == "/games/"
}

// RespondTo implement interfaces.Endpoint
func (ep Games) RespondTo(r *http.Request) (interfaces.Entity, error) {
	return entities.PostGames{ID: uuid.Generate()}, nil
}

// Validate interfaces.Endpoint
func (ep Games) Validate(r *http.Request) (interfaces.Entity, error) {
	if r.Method != "POST" {
		return entities.MethodNotAllowed{}, nil
	}

	return nil, nil
}
