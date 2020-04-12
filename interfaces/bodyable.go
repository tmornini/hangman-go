package interfaces

import "io"

type Bodyable interface {
	Getable
	WriteBodyTo(io.Writer)
}
