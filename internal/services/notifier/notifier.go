package notifier

import (
	"context"
	"errors"
	"github.com/AnKlvy/notifier-edunite/internal/database"
	"github.com/nikoksr/notify"
)

type NotifyInterface interface {
	notify.Notifier
	AddReceivers(receivers ...string)
}

func NewNotifyService(repo database.Models, services map[string]NotifyInterface) *NotifyService {
	return &NotifyService{repo: repo, services: services}
}

type NotifyService struct {
	repo     database.Models
	services map[string]NotifyInterface
}

func (n *NotifyService) Subscribe(userId string, channel string, token string) error {

	err := n.repo.Notifier.Subscribe(userId, channel, token)
	if err != nil {
		return err
	}
	return nil
}

func (n *NotifyService) Unsubscribe(userId string, channel string) error {

	err := n.repo.Notifier.Unsubscribe(userId, channel)
	if err != nil {
		return err
	}
	return nil
}

func (n *NotifyService) SendNotification(ctx context.Context, userId string, notification *database.Notification) error {

	for channel, service := range n.services {
		token, err := n.repo.Notifier.GetReceiverByUserAndChannel(userId, channel)
		if err != nil {
			if errors.Is(err, database.ErrRecordNotFound) {
				continue // пользователь не подписан на этот канал — пропускаем
			}
			return err // другая ошибка — возвращаем
		}

		service.AddReceivers(*token)

		if &token != nil {
			err = service.Send(ctx, notification.Subject, notification.Message)
			if err != nil {
				return err // можно также логировать и продолжить, если не критично
			}
		}
	}
	err := n.repo.Notifier.SendNotification(userId, notification)
	if err != nil {
		return err
	}
	return nil
}
