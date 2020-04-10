package interfaces

type StorableEntity interface {
	Entity
	MetaData() Metadata
}
