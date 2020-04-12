package responses

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/interfaces"
)

type MethodNotAllowed struct {
	interfaces.Getable
}

// SerializeTo implement interfaces.Responsible
func (res MethodNotAllowed) SerializeTo(w http.ResponseWriter) error {
	w.WriteHeader(405)

	return nil
}
