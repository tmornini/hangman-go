package authorizers

import "github.com/tmornini/udemy-hangman/interfaces"

type Public struct {
}

func (authn Public) Authorize(req interfaces.Requestable) interfaces.Responsible {
	return nil
}
