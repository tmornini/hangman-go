package interfaces

type Authorizable interface {
	Authorize(Requestable) Responsible
}
