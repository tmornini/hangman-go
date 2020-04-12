package interfaces

type Pathable interface {
	PathPrefix() string
	PathPostfix() string
}
