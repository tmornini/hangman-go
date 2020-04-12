package responses

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/interfaces"
)

type ServiceNotAvailable struct {
	interfaces.Getable `json:"-"`
}

// SerializeTo implement interfaces.Responsible
func (ety ServiceNotAvailable) SerializeTo(w http.ResponseWriter) error {
	w.WriteHeader(503)

	return nil
}
