package authorizers

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/interfaces"
	"github.com/tmornini/udemy-hangman/responses"
)

type Token struct {
}

func (authn Token) Authorize(r *http.Request) interfaces.Responsible {
	authorization := r.Header.Get("Authorization")

	if authorization == "Bearer server-token" {
		return nil
	}

	return responses.Unauthorized{}
}
