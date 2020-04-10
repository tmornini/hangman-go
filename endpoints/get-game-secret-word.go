package endpoints

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/interfaces"
	"github.com/tmornini/udemy-hangman/responses"
	"github.com/tmornini/udemy-hangman/secretwords"
)

type GetGameSecretWord struct{}

// RespondsToPathOf implement interfaces.Endpoint
func (ep GetGameSecretWord) RespondsToPathOf(r *http.Request) bool {
	return r.URL.Path == "/games/"
}

// RespondsToMethodOf interfaces.Endpoint
func (ep GetGameSecretWord) RespondsToMethodOf(r *http.Request) bool {
	return r.Method == "GET"
}

// Process implement interfaces.Endpoint
func (ep GetGameSecretWord) Process(r *http.Request) (interfaces.Responsible, error) {
	return responses.GetGameSecretWord{SecretWord: secretwords.ChooseOne()}, nil
}
