package interfaces

import "net/http"

type Endpointable interface {
	Authorizable
	RespondsToPathOf(*http.Request) bool
	RespondsToMethodOf(*http.Request) bool
	Process(*http.Request) (Responsible, error)
}
