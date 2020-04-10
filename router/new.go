package router

import "github.com/tmornini/udemy-hangman/interfaces"

func New(endpoints ...interfaces.Endpointable) Router {
	return Router{
		endpoints: endpoints,
	}
}
