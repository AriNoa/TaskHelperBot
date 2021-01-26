package echo

// Interactor is a struct that implements Handler interface
type Interactor struct {
	Presenter Presenter
}

// Handle is a method that implements Handler's Handle method
func (i *Interactor) Handle(contents string) {
	i.Presenter.echo(contents)
}
