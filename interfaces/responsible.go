package interfaces

import "net/http"

type Responsible interface {
	WriteTo(http.ResponseWriter) error
}
