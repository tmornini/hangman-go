package interfaces

type Store interface {
	Get(Metadata) (StorableEntity, error)
	GetCollection(Metadata) ([]StorableEntity, error)
	GetHistory(Metadata) ([]StorableEntity, error)
	Put(StorableEntity) error
	PutIfMatch(StorableEntity) error
}
