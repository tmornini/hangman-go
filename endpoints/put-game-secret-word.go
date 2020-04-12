package endpoints

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/tmornini/udemy-hangman/interfaces"
	"github.com/tmornini/udemy-hangman/responses"
)

type PutGameSecretWord struct {
	interfaces.Authorizable
}

// RespondsToPathOf implement interfaces.Endpoint
func (ep PutGameSecretWord) RespondsToPathOf(r *http.Request) bool {
	_, err := ep.gameIDFrom(r)

	return err == nil
}

// RespondsToMethodOf implement interfaces.Endpoint
func (ep PutGameSecretWord) RespondsToMethodOf(r *http.Request) bool {
	return r.Method == "PUT"
}

// Process implement interfaces.Endpoint
func (ep PutGameSecretWord) Process(r *http.Request) (interfaces.Responsible, error) {
	return responses.PutGameSecretWord{
		SecretWord: "secret",
	}, nil
}

var gameIDRegexp = regexp.MustCompile("/games/([^/]+)/secret-word")

func (ep PutGameSecretWord) gameIDFrom(r *http.Request) (string, error) {

	matches := gameIDRegexp.FindStringSubmatch(r.URL.Path)

	if len(matches) == 0 {
		return "", errors.New("could not match path")
	}

	return matches[1], nil
}
