package authorizers

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/interfaces"
)

type Public struct {
}

func (authn Public) Authorize(r *http.Request) interfaces.Responsible {
	return nil
}
