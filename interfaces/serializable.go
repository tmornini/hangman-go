package interfaces

import "net/http"

type Serializable interface {
	Serialize(*http.Request, Responsible) []byte
}
