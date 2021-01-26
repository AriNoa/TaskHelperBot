package echo

// Interactor is a struct that implements UseCase interface
type Interactor struct {
	Presenter Presenter
}

// Echo is a method that implements UseCase's Echo method
func (i *Interactor) Echo(contents string) {
	i.Presenter.Echo(contents)
}
