package interfaces

type Metadata interface {
	Prefix() string
	ID() string
	ETag() string
	CreatedAt() string
}
