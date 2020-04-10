package interfaces

type StorableResponse interface {
	Responsible
	MetaData() Metadata
}
