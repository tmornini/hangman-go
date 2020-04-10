package router

import "github.com/tmornini/udemy-hangman/interfaces"

type Router struct {
	endpoints []interfaces.Endpointable
}
