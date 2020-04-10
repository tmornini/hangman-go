package endpoints

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/interfaces"
	"github.com/tmornini/udemy-hangman/responses"
)

type GetErrorAnticipated struct{}

// RespondsToURLOf interfaces.Endpointable
func (ep GetErrorAnticipated) RespondsToPathOf(r *http.Request) bool {
	return r.URL.Path == "/error-anticipated"
}

// RespondsToMethodOf interfaces.Endpointable
func (ep GetErrorAnticipated) RespondsToMethodOf(r *http.Request) bool {
	return r.Method == "GET"
}

// Process interfaces.Endpointable
func (ep GetErrorAnticipated) Process(r *http.Request) (interfaces.Responsible, error) {
	return responses.ServiceNotAvailable{}, nil
}
