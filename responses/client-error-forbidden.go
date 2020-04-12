package responses

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/interfaces"
)

type Forbidden struct {
	interfaces.Getable
}

// SerializeTo implement interfaces.Responsible
func (res Forbidden) SerializeTo(w http.ResponseWriter) error {
	w.Header().Add("WWW-Authenticate", "Bearer realm=hangman")
	w.WriteHeader(403)

	return nil
}
