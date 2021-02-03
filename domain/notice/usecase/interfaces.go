package usecase

import (
	"time"

	model "github.com/AriNoa/TaskHelperBot/domain/notice/model"
)

// UseCase is an interface for notice
type UseCase interface {
	CreateTable(userID string, tableID string)
	DeleteTable(userID string, tableID string)
	CopyTable(userID string, tableID string, sourceID string)

	AppendNotice(userID string, tableID string, contents string, time time.Time)
	RemoveNotice(userID string, tableID string, noticeID int)

	Repeat(userID string, tableID string, noticeID int, interval time.Duration)

	UserOn(userID string)
	TableOn(userID string, tableID string)
	NoticeOn(userID string, tableID string, noticeID int)

	UserOf(userID string)
	TableOff(userID string, tableID string)
	NoticeOff(userID string, tableID string, noticeID int)

	TableList(userID string)
	NoticeList(userID string, tableID string)
	Info(userID string, tableID string, noticeID int)
}

// Presenter is an interface to tell the notice user
type Presenter interface {
	Complete(event string)
	Error(event string, contents string)

	DrawTableList(user model.User)
	DrawNoticeList(table model.Table)
	DrawNoticeInfo(notice model.Notice)
}
