package interfaces

import "net/http"

type Entity interface {
	WriteResponseTo(http.ResponseWriter) error
}
