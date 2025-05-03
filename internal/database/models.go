package database

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	// Устанавливаем поле News как интерфейс, содержащий методы, которые должны поддерживать
	// как 'реальная' модель, так и мок-модель.
	Notifier interface {
		Subscribe(userId string, channel string, token string) error
		Unsubscribe(userId string, channel string) error
		SendNotification(userId string, notification *Notification) error
		GetReceiverByUserAndChannel(userId string, channel string) (*string, error)
	}
}

func NewNotifier(db *sql.DB) Models {
	return Models{
		Notifier: &NotifierModel{DB: db},
	}
}
