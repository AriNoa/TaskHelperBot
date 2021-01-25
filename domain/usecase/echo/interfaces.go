package echo

// Handler is an interface for echo
type Handler interface {
	Handle(contents string)
}

// Presenter is an interface for echo the same contents
type Presenter interface {
	echo(contents string)
}
