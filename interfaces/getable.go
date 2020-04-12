package interfaces

type Getable interface {
	Get(string) (bool, []string)
}
