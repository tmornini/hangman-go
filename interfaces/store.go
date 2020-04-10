package interfaces

type Store interface {
	Get(Metadata) (StorableResponse, error)
	GetCollection(Metadata) ([]StorableResponse, error)
	GetHistory(Metadata) ([]StorableResponse, error)
	Put(StorableResponse) error
	PutIfMatch(StorableResponse) error
}
