package endpoints

import (
	"net/http"
	"regexp"

	"github.com/tmornini/udemy-hangman/interfaces"
	"github.com/tmornini/udemy-hangman/responses"
)

type PutGameSecretWord struct {
	id string
}

// RespondsToPathOf implement interfaces.Endpoint
func (ep PutGameSecretWord) RespondsToPathOf(r *http.Request) bool {
	re := regexp.MustCompile("/games/([^/]+)/secret-word")

	matches := re.FindStringSubmatch(r.URL.Path)

	if len(matches) == 0 {
		return false
	}

	ep.id = matches[1]

	return true
}

// RespondsToMethodOf implement interfaces.Endpoint
func (ep PutGameSecretWord) RespondsToMethodOf(r *http.Request) bool {
	return r.Method == "PUT"
}

// Process implement interfaces.Endpoint
func (ep PutGameSecretWord) Process(r *http.Request) (interfaces.Responsible, error) {
	if r.Header.Get("Authorization") != "server-token" {
		return responses.Unauthorized{}, nil
	}

	return responses.PutGameSecretWord{
		ID:         ep.id,
		SecretWord: "secret",
	}, nil
}
