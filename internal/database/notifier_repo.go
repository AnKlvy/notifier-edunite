package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/AnKlvy/notifier-edunite/internal/validator"
	"time"
)

type Notification struct {
	Message  string
	Subject  string
	Images   *[]string
	Metadata map[string]string
}

type NotifierSettings struct {
	UserId  string
	Channel string
	Token   string
}

// ValidateNews выполняет валидацию данных новости.
func ValidateNotification(v *validator.Validator, notification *Notification) {
	v.Check(notification.Message != "", "message", "must be provided")
	v.Check(notification.Subject != "", "subject", "must be provided")
}

func ValidateSettings(v *validator.Validator, userId, channel, token string) {
	v.Check(userId != "", "user_id", "must be provided")
	v.Check(token != "", "token", "must be provided")
	v.Check(channel != "", "channel", "must be provided")
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
func (n *NotifierModel) SendNotification(userId string, notification *Notification) error {
	// Преобразуем map в JSON
	metadataJSON, err := json.Marshal(notification.Metadata)
	if err != nil {
		return fmt.Errorf("failed to marshal metadata: %w", err)
	}

	query := `
   INSERT INTO notifications (user_id, message, subject, metadata)
   VALUES ($1, $2, $3, $4)
   RETURNING id`
	args := []any{userId, notification.Message, notification.Subject, metadataJSON}

	// Создаём контекст с тайм-аутом 3 секунды.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err = n.DB.ExecContext(ctx, query, args...)
	// Используем QueryRowContext() и передаём контекст в качестве первого аргумента.
	return err
}

func (n *NotifierModel) GetReceiversByUserAndChannel(userId, channel string) ([]string, error) {
	query := `
		SELECT token FROM notifier_settings
		WHERE user_id = $1 AND channel = $2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := n.DB.QueryContext(ctx, query, userId, channel)
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
