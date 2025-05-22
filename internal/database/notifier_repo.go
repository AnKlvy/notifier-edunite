package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/AnKlvy/notifier-edunite/internal/validator"
	"github.com/lib/pq"
	"time"
)

type Notification struct {
	Id        int
	Message   string
	Subject   string
	Images    *[]string
	Metadata  map[string]string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type NotifierSettings struct {
	Id      int
	UserId  string
	Channel string
	Token   string
}

// ValidateNews выполняет валидацию данных новости.

func ValidateSettings(v *validator.Validator, userId, channel string) {
	v.Check(userId != "", "user_id", "must be provided")
	v.Check(len(userId) <= 100, "user_id", "must be no more than 100 characters")
	v.Check(validator.PermittedValue(channel, "email", "firebase"), "channel", "invalid channel")
}

func ValidateSubscribe(v *validator.Validator, userId, channel, token string) {
	ValidateSettings(v, userId, channel)

	if channel == "email" {
		v.Check(validator.Matches(token, validator.EmailRX), "token", "must be a valid email address")
	}
}

func ValidateNotification(v *validator.Validator, notification Notification) {
	v.Check(notification.Message != "", "message", "message is required")
	v.Check(len(notification.Message) <= 2000, "message", "message must be no more than 2000 characters")

	v.Check(notification.Subject != "", "subject", "subject is required")
	v.Check(len(notification.Subject) <= 255, "subject", "subject must be no more than 255 characters")

	for i, img := range *notification.Images {
		v.Check(img != "", "images", "image URL cannot be empty")
		v.Check(len(img) <= 1000, "images", "image URL is too long")
		v.Check(i < 10, "images", "too many images (max 10)")
	}
}

type NotifierModel struct {
	DB *sql.DB
}

func (n *NotifierModel) Subscribe(userId, channel, token string) error {
	query := `
   INSERT INTO notifier_settings (user_id, channel, token)
   VALUES ($1, $2, $3)
   RETURNING id`
	args := []any{userId, channel, token}

	// Создаём контекст с тайм-аутом 3 секунды.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := n.DB.ExecContext(ctx, query, args...)
	// Используем QueryRowContext() и передаём контекст в качестве первого аргумента.
	return err
}

func (n *NotifierModel) Unsubscribe(userId, channel string) error {
	query := `
    DELETE FROM notifier_settings
    WHERE user_id = $1 and channel = $2`

	args := []any{userId, channel}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := n.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

func (n *NotifierModel) SendNotification(userIds []string, notification *Notification) error {
	metadataJSON, err := json.Marshal(notification.Metadata)
	if err != nil {
		return fmt.Errorf("failed to marshal metadata: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := n.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var notificationId int
	now := time.Now()
	query := `INSERT INTO notifications (message, subject, metadata, images, created_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err = tx.QueryRowContext(ctx, query, notification.Message, notification.Subject, metadataJSON,
		pq.Array(notification.Images), now, now).Scan(&notificationId)
	if err != nil {
		return err
	}

	// Обновляем время в структуре
	notification.CreatedAt = now
	notification.UpdatedAt = now

	insertUserQuery := `INSERT INTO notifications_users (notification_id, user_id) VALUES ($1, $2)`
	for _, userId := range userIds {
		_, err := tx.ExecContext(ctx, insertUserQuery, notificationId, userId)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (n *NotifierModel) GetReceiversByUsersAndChannel(userIds []string, channel string) ([]string, error) {
	query := `
		SELECT token FROM notifier_settings
		WHERE user_id = ANY($1) AND channel = $2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := n.DB.QueryContext(ctx, query, pq.Array(userIds), channel)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return n.getTokensByRows(rows)
}

func (n *NotifierModel) GetAllReceiversByChannel(channel string) ([]string, error) {
	query := `
		SELECT token FROM notifier_settings
		WHERE channel = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := n.DB.QueryContext(ctx, query, channel)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return n.getTokensByRows(rows)
}

func (n *NotifierModel) getTokensByRows(rows *sql.Rows) ([]string, error) {
	var tokens []string
	for rows.Next() {
		var token string
		if err := rows.Scan(&token); err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(tokens) == 0 {
		return nil, ErrRecordNotFound
	}

	return tokens, nil
}

// Получение всех настроек уведомлений
func (n *NotifierModel) GetAllSettings() ([]NotifierSettings, error) {
	query := `
		SELECT * FROM notifier_settings `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := n.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var settings []NotifierSettings
	for rows.Next() {
		var s NotifierSettings
		if err := rows.Scan(&s.Id, &s.UserId, &s.Channel, &s.Token); err != nil {
			return nil, err
		}
		settings = append(settings, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return settings, nil
}

// Получение всех уведомлений
func (n *NotifierModel) GetAllNotifications() ([]Notification, error) {
	query := `
		SELECT *
		FROM notifications
		ORDER BY created_at DESC`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := n.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []Notification
	for rows.Next() {
		var n Notification
		var metadataJSON []byte
		var images []string

		if err := rows.Scan(&n.Id, &n.Message, &n.Subject, &metadataJSON, pq.Array(&images), &n.CreatedAt, &n.UpdatedAt); err != nil {
			return nil, err
		}

		// Преобразуем JSON в map
		if len(metadataJSON) > 0 {
			if err := json.Unmarshal(metadataJSON, &n.Metadata); err != nil {
				return nil, err
			}
		}

		n.Images = &images
		notifications = append(notifications, n)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notifications, nil
}
