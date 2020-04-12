package interfaces

import "io"

type Headerable interface {
	Getable
	WriteHeaderTo(io.Writer) (int, error)
}
