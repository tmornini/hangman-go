package responses

import "net/http"

type ServiceNotAvailable struct{}

// WriteTo implement interfaces.Responsible
func (ety ServiceNotAvailable) WriteTo(w http.ResponseWriter) error {
	w.WriteHeader(503)

	return nil
}
