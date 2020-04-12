package authorizers

import (
	"github.com/tmornini/udemy-hangman/interfaces"
	"github.com/tmornini/udemy-hangman/responses"
)

type Token struct {
}

func (authn Token) Authorize(req interfaces.Requestable) interfaces.Responsible {
	exists, authorizations := req.Get("header.Authorization")

	if exists && authorizations[0] == "Bearer server-token" {
		return nil
	}

	if exists {
		return responses.Forbidden{}
	}

	return responses.Unauthorized{}
}
