package usecase

import (
	model "github.com/AriNoa/TaskHelperBot/domain/notice/model"
)

type Interactor struct {
	Presenter  Presenter
	Repository model.UserRepository
}

func (i *Interactor) CreateTable(userID string, tableID string) {

}
