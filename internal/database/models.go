package database

import (
	"database/sql"
	"errors"
	"time"
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
		SendNotification(userId []string, notification *Notification) error
		GetReceiversByUsersAndChannel(userId []string, channel string) ([]string, error)
		GetAllReceiversByChannel(channel string) ([]string, error)
		GetAllSettings() ([]NotifierSettings, error)
		GetAllNotifications() ([]Notification, error)
		GetUserSettings(userId string) ([]NotifierSettings, error)
		GetUserNotifications(userId string, fromDate time.Time) ([]Notification, error)
	}
}

func NewNotifier(db *sql.DB) Models {
	return Models{
		Notifier: &NotifierModel{DB: db},
	}
}
