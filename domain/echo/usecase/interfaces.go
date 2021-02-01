package echo

// UseCase is an interface for echo
type UseCase interface {
	Echo(contents string)
}

// Presenter is an interface for echo the same contents
type Presenter interface {
	Echo(contents string)
}
