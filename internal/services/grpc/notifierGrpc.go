package grpc

import (
	"context"
	"errors"
	"github.com/AnKlvy/notifier-edunite/internal/database"
	"github.com/AnKlvy/notifier-edunite/internal/services/notifier"
	"github.com/AnKlvy/notifier-edunite/internal/validator"
	"github.com/AnKlvy/notifier-edunite/protobuf/gen_notifier"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
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
	v.Check(request.GetValue() != "", "token", "must be provided")
	v.Check(len(request.GetValue()) <= 500, "token", "must be no more than 500 characters")
	database.ValidateSubscribe(v, request.GetUserId(), request.GetChannel(), request.GetValue())

	if !v.Valid() {
		// Преобразуем карту ошибок в строку
		errorMessages := ""
		for field, message := range v.Errors {
			if errorMessages != "" {
				errorMessages += "; "
			}
			errorMessages += field + ": " + message
		}

		return &gen_notifier.SuccessResponse{
			Success:      false,
			ErrorMessage: errorMessages,
		}, errors.New(errorMessages)
	}

	err := s.notifySrv.Subscribe(request.GetUserId(), request.GetChannel(), request.GetValue())
	if err != nil {
		return &gen_notifier.SuccessResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, err
	}

	return &gen_notifier.SuccessResponse{Success: true, ErrorMessage: ""}, nil
}

func (s *Service) Unsubscribe(ctx context.Context, request *gen_notifier.UnsubscribeRequest) (*gen_notifier.SuccessResponse, error) {
	v := validator.New()
	database.ValidateSettings(v, request.GetUserId(), request.GetChannel())

	if !v.Valid() {
		return &gen_notifier.SuccessResponse{
			Success:      false,
			ErrorMessage: "invalid input data",
		}, errors.New("invalid input data")
	}

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

	v := validator.New()
	database.ValidateNotification(v, notifi)

	if !v.Valid() {
		return nil, errors.New("invalid notification input")
	}

	err := s.notifySrv.SendToOneOrManyByChannel(ctx, notification.GetUsersIds(), &notifi)
	if err != nil {
		return nil, err
	}

	// Обновляем поля в ответе
	if notification.Notification != nil {
		notification.Notification.Id = int64(notifi.Id)
		notification.Notification.CreatedAt = notifi.CreatedAt.Format(time.RFC3339)
		notification.Notification.UpdatedAt = notifi.UpdatedAt.Format(time.RFC3339)
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

	v := validator.New()

	database.ValidateNotification(v, notifi)

	if !v.Valid() {
		return nil, errors.New("invalid notification input")
	}

	err := s.notifySrv.SendToAll(ctx, &notifi)
	if err != nil {
		return nil, err
	}

	// Обновляем поля в ответе
	notification.Id = int64(notifi.Id)
	notification.CreatedAt = notifi.CreatedAt.Format(time.RFC3339)
	notification.UpdatedAt = notifi.UpdatedAt.Format(time.RFC3339)

	return notification, nil
}

// Получение всех настроек уведомлений
func (s *Service) GetAllSettings(ctx context.Context, empty *emptypb.Empty) (*gen_notifier.GetAllSettingsResponse, error) {
	settings, err := s.notifySrv.GetAllSettings()
	if err != nil {
		return nil, err
	}

	response := &gen_notifier.GetAllSettingsResponse{
		Settings: make([]*gen_notifier.NotifierSettings, 0, len(settings)),
	}

	for _, setting := range settings {
		response.Settings = append(response.Settings, &gen_notifier.NotifierSettings{
			UserId:  setting.UserId,
			Channel: setting.Channel,
			Token:   setting.Token,
		})
	}

	return response, nil
}

// Получение всех уведомлений
func (s *Service) GetAllNotifications(ctx context.Context, empty *emptypb.Empty) (*gen_notifier.GetAllNotificationsResponse, error) {
	notifications, err := s.notifySrv.GetAllNotifications()
	if err != nil {
		return nil, err
	}

	response := &gen_notifier.GetAllNotificationsResponse{
		Notifications: make([]*gen_notifier.Notification, 0, len(notifications)),
	}

	for _, notification := range notifications {
		var images []string
		if notification.Images != nil {
			images = *notification.Images
		}

		response.Notifications = append(response.Notifications, &gen_notifier.Notification{
			Id:        int64(notification.Id), // Добавляем ID из базы данных
			Message:   notification.Message,
			Subject:   notification.Subject,
			Images:    images,
			Metadata:  notification.Metadata,
			CreatedAt: notification.CreatedAt.Format(time.RFC3339),
			UpdatedAt: notification.UpdatedAt.Format(time.RFC3339),
		})
	}

	return response, nil
}

// Получение настроек конкретного пользователя
func (s *Service) GetUserSettings(ctx context.Context, request *gen_notifier.GetUserSettingsRequest) (*gen_notifier.GetAllSettingsResponse, error) {
	v := validator.New()
	v.Check(request.GetUserId() != "", "user_id", "must be provided")

	if !v.Valid() {
		return nil, errors.New("invalid user_id")
	}

	settings, err := s.notifySrv.GetUserSettings(request.GetUserId())
	if err != nil {
		return nil, err
	}

	response := &gen_notifier.GetAllSettingsResponse{
		Settings: make([]*gen_notifier.NotifierSettings, 0, len(settings)),
	}

	for _, setting := range settings {
		response.Settings = append(response.Settings, &gen_notifier.NotifierSettings{
			UserId:  setting.UserId,
			Channel: setting.Channel,
			Token:   setting.Token,
		})
	}

	return response, nil
}

// Получение уведомлений конкретного пользователя
func (s *Service) GetUserNotifications(ctx context.Context, request *gen_notifier.GetUserNotificationsRequest) (*gen_notifier.GetAllNotificationsResponse, error) {
	v := validator.New()
	v.Check(request.GetUserId() != "", "user_id", "must be provided")
	v.Check(len(request.GetUserId()) <= 100, "user_id", "must be no more than 100 characters")

	if !v.Valid() {
		return nil, errors.New("invalid user_id")
	}

	// Парсим дату из строки
	var fromDate time.Time
	var err error
	if request.GetFromDate() != "" {
		fromDate, err = time.Parse(time.RFC3339, request.GetFromDate())
		if err != nil {
			return nil, errors.New("invalid date format, use ISO 8601 (RFC3339)")
		}
	} else {
		// Если дата не указана, используем дату "начала времен"
		fromDate = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	}

	notifications, err := s.notifySrv.GetUserNotifications(request.GetUserId(), fromDate)
	if err != nil {
		return nil, err
	}

	response := &gen_notifier.GetAllNotificationsResponse{
		Notifications: make([]*gen_notifier.Notification, 0, len(notifications)),
	}

	for _, notification := range notifications {
		var images []string
		if notification.Images != nil {
			images = *notification.Images
		}

		response.Notifications = append(response.Notifications, &gen_notifier.Notification{
			Id:        int64(notification.Id),
			Message:   notification.Message,
			Subject:   notification.Subject,
			Images:    images,
			Metadata:  notification.Metadata,
			CreatedAt: notification.CreatedAt.Format(time.RFC3339),
			UpdatedAt: notification.UpdatedAt.Format(time.RFC3339),
		})
	}

	return response, nil
}
