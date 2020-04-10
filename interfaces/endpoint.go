package interfaces

import "net/http"

type Endpoint interface {
	RespondsToPathOf(*http.Request) bool
	RespondTo(*http.Request) (Entity, error)
	Validate(*http.Request) (Entity, error)
}
