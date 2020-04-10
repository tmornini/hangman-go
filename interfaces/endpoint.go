package interfaces

import "net/http"

type Endpoint interface {
	RespondsTo(*http.Request) bool
	RespondTo(*http.Request) (Entity, error)
}
