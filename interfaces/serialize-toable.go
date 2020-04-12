package interfaces

import "net/http"

type SerializeToable interface {
	SerializeTo(http.ResponseWriter) error
}
