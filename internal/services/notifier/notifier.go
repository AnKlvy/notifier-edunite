package notifier

import (
	"context"
	"errors"
	"fmt"
	"github.com/AnKlvy/notifier-edunite/internal/database"
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

func (n *NotifyService) SendToOneOrManyByChannel(ctx context.Context, userIds []string, notification *database.Notification) error {
	for channel, service := range n.services {
		tokens, err := n.repo.Notifier.GetReceiversByUsersAndChannel(userIds, channel)
		err = n.send(ctx, service, tokens, err, notification)
		if err != nil {
			return err
		}
	}
	err := n.repo.Notifier.SendNotification(userIds, notification)
	if err != nil {
		return err
	}
	return nil
}

func (n *NotifyService) SendToAll(ctx context.Context, notification *database.Notification) error {
	for channel, service := range n.services {
		tokens, err := n.repo.Notifier.GetAllReceiversByChannel(channel)
		err = n.send(ctx, service, tokens, err, notification)
	}
	err := n.repo.Notifier.SendNotification([]string{"all"}, notification)
	if err != nil {
		return err
	}
	return nil
}

func (n *NotifyService) send(ctx context.Context, service NotifyInterface, tokens []string, err error, notification *database.Notification) error {
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			return nil // пользователь не подписан на этот канал — пропускаем
		}
		return err // другая ошибка — возвращаем
	}
	err = service.Send(ctx, notification.Subject, notification.Message, tokens, *notification.Images...)
	if err != nil {
		return fmt.Errorf("error sending message to tokens: %v", err)
	}
	return nil
}
