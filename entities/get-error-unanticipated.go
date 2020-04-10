package entities

import "net/http"

type GetErrorUnanticipatedEntity struct{}

// WriteResponseTo implement interfaces.Entity
func (ety GetErrorUnanticipatedEntity) WriteResponseTo(w http.ResponseWriter) error {
	w.WriteHeader(200)

	return nil
}
