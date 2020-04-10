package endpoints

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/entities"
	"github.com/tmornini/udemy-hangman/interfaces"
	"github.com/tmornini/udemy-hangman/secretwords"
)

type GameSecretWord struct{}

// RespondsToPathOf implement interfaces.Endpoint
func (ep GameSecretWord) RespondsToPathOf(r *http.Request) bool {
	return r.URL.Path == "/games/"
}

// RespondTo implement interfaces.Endpoint
func (ep GameSecretWord) RespondTo(r *http.Request) (interfaces.Entity, error) {
	return entities.GetGameSecretWord{SecretWord: secretwords.ChooseOne()}, nil
}

// Validate interfaces.Endpoint
func (ep GameSecretWord) Validate(r *http.Request) (interfaces.Entity, error) {
	if r.Method != "PUT" && r.Method != "GET" {
		return entities.MethodNotAllowed{}, nil
	}

	return nil, nil
}
