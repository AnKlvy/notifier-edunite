package grpc

import (
	"context"
	"errors"
	"github.com/AnKlvy/notifier-edunite/internal/database"
	"github.com/AnKlvy/notifier-edunite/internal/services/notifier"
	"github.com/AnKlvy/notifier-edunite/internal/validator"
	"github.com/AnKlvy/notifier-edunite/protobuf/gen_notifier"
	"google.golang.org/grpc"
)

type Service struct {
	notifySrv notifier.NotifyService
	gen_notifier.UnimplementedNotificationServiceServer
}

func NewNotifierService(grpc *grpc.Server, notifySrv notifier.NotifyService) {
	notifierService := &Service{notifySrv: notifySrv}
	gen_notifier.RegisterNotificationServiceServer(grpc, notifierService)
}
func (s *Service) Subscribe(ctx context.Context, request *gen_notifier.SubscribeRequest) (*gen_notifier.SuccessResponse, error) {
	v := validator.New()
	if database.ValidateSettings(v, request.GetUserId(), request.GetChannel(), request.GetValue()); !v.Valid() {
		return nil, errors.New("invalid settings input data")
	}

	err := s.notifySrv.Subscribe(request.GetUserId(), request.GetChannel(), request.GetValue())
	if err != nil {
		return nil, err
	}

	return &gen_notifier.SuccessResponse{Success: true, ErrorMessage: ""}, nil
}

func (s *Service) Unsubscribe(ctx context.Context, request *gen_notifier.UnsubscribeRequest) (*gen_notifier.SuccessResponse, error) {
	err := s.notifySrv.Unsubscribe(request.GetUserId(), request.GetChannel())
	if err != nil {
		return &gen_notifier.SuccessResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, err
	}
	return &gen_notifier.SuccessResponse{Success: true, ErrorMessage: ""}, nil
}

func (s *Service) SendToOneOrMany(ctx context.Context, notification *gen_notifier.UsersNotification) (*gen_notifier.UsersNotification, error) {
	images := notification.GetNotification().GetImages()
	notifi := database.Notification{
		Message:  notification.GetNotification().GetMessage(),
		Subject:  notification.GetNotification().GetSubject(),
		Images:   &images,
		Metadata: notification.GetNotification().GetMetadata(),
	}

	err := s.notifySrv.SendToOneOrManyByChannel(ctx, notification.GetUsersIds(), &notifi)
	if err != nil {
		return nil, err
	}

	return notification, nil
}

func (s *Service) SendToAll(ctx context.Context, notification *gen_notifier.Notification) (*gen_notifier.Notification, error) {
	images := notification.GetImages()
	notifi := database.Notification{
		Message:  notification.GetMessage(),
		Subject:  notification.GetSubject(),
		Images:   &images,
		Metadata: notification.GetMetadata(),
	}

	err := s.notifySrv.SendToAll(ctx, &notifi)
	if err != nil {
		return nil, err
	}

	return notification, nil
}
