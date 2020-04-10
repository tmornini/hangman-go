package endpoints

import (
	"errors"
	"net/http"

	"github.com/tmornini/udemy-hangman/interfaces"
	"github.com/tmornini/udemy-hangman/responses"
)

type GetErrorUnanticipated struct{}

// RespondsToPathOf implement interfaces.Endpoint
func (ep GetErrorUnanticipated) RespondsToPathOf(r *http.Request) bool {
	return r.URL.Path == "/error-unanticipated"
}

// RespondsToMethodOf implement interfaces.Endpoint

func (ep GetErrorUnanticipated) RespondsToMethodOf(r *http.Request) bool {
	return r.Method == "GET"
}

// Process implement interfaces.Endpoint
func (ep GetErrorUnanticipated) Process(r *http.Request) (interfaces.Responsible, error) {
	if true {
		return nil, errors.New("this error was not anticipated!")
	}

	return responses.ServiceNotAvailable{}, nil
}
