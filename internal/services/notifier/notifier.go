package notifier

import (
	"context"
	"errors"
	"fmt"
	"github.com/AnKlvy/notifier-edunite/internal/database"
	"log"
	"time"
)

type NotifyInterface interface {
	Send(ctx context.Context, subject string, message string, receivers []string, images ...string) error
}

func NewNotifyService(repo database.Models, services map[string]NotifyInterface) *NotifyService {
	return &NotifyService{repo: repo, services: services}
}

type NotifyService struct {
	repo     database.Models
	services map[string]NotifyInterface
}

func (n *NotifyService) Subscribe(userId string, channel string, token string) error {
	return n.repo.Notifier.Subscribe(userId, channel, token)
}

func (n *NotifyService) Unsubscribe(userId string, channel string) error {
	return n.repo.Notifier.Unsubscribe(userId, channel)
}

func (n *NotifyService) SendToOneOrManyByChannel(ctx context.Context, userIds []string, notification *database.Notification) error {
	for channel, service := range n.services {
		tokens, err := n.repo.Notifier.GetReceiversByUsersAndChannel(userIds, channel)
		if sendErr := n.send(ctx, service, tokens, err, notification); sendErr != nil {
			log.Printf("SendToOneOrManyByChannel: channel %s — %v", channel, sendErr)
		}
	}

	if err := n.repo.Notifier.SendNotification(userIds, notification); err != nil {
		log.Printf("SendToOneOrManyByChannel: db log error — %v", err)
	}
	return nil
}

func (n *NotifyService) SendToAll(ctx context.Context, notification *database.Notification) error {
	for channel, service := range n.services {
		tokens, err := n.repo.Notifier.GetAllReceiversByChannel(channel)
		if sendErr := n.send(ctx, service, tokens, err, notification); sendErr != nil {
			log.Printf("SendToAll: channel %s — %v", channel, sendErr)
		}
	}

	if err := n.repo.Notifier.SendNotification([]string{"all"}, notification); err != nil {
		log.Printf("SendToAll: db log error — %v", err)
	}
	return nil
}

func (n *NotifyService) send(ctx context.Context, service NotifyInterface, tokens []string, err error, notification *database.Notification) error {
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return nil // канал просто не подписан — это не ошибка
		}
		return err
	}

	err = service.Send(ctx, notification.Subject, notification.Message, tokens, *notification.Images...)
	if err != nil {
		return fmt.Errorf("error sending message to tokens: %v", err)
	}
	return nil
}

// Получение всех настроек уведомлений
func (n *NotifyService) GetAllSettings() ([]database.NotifierSettings, error) {
	return n.repo.Notifier.GetAllSettings()
}

// Получение всех уведомлений
func (n *NotifyService) GetAllNotifications() ([]database.Notification, error) {
	return n.repo.Notifier.GetAllNotifications()
}

// Получение настроек конкретного пользователя
func (n *NotifyService) GetUserSettings(userId string) ([]database.NotifierSettings, error) {
	return n.repo.Notifier.GetUserSettings(userId)
}

// Получение уведомлений по пользователю с учетом даты
func (n *NotifyService) GetUserNotifications(userId string, fromDate time.Time) ([]database.Notification, error) {
	return n.repo.Notifier.GetUserNotifications(userId, fromDate)
}
