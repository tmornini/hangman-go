package interfaces

type Storagable interface {
	Get(Requestable) (Responsible, error)
	GetCollection(Requestable) ([]Responsible, error)
	GetHistory(Requestable) ([]Responsible, error)
	Put(Responsible) error
	PutIfMatch(Responsible) error
}
