package interfaces

import "net/http"

type Authorizable interface {
	Authorize(*http.Request) Responsible
}
