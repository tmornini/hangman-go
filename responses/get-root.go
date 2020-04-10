package responses

import "net/http"

type GetRoot struct{}

// WriteTo implement interfaces.Responsible
func (ety GetRoot) WriteTo(w http.ResponseWriter) error {
	w.WriteHeader(200)

	return nil
}
