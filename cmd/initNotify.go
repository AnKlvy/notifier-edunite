package main

import (
	"context"
	"database/sql"
	"github.com/AnKlvy/notifier-edunite/internal/database"
	"github.com/AnKlvy/notifier-edunite/internal/services/notifier"
	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/fcm"
	"github.com/nikoksr/notify/service/mail"
)

func initNotify(db *sql.DB) (*notifier.NotifyService, error) {
	//init repo interface
	repo := database.NewNotifier(db)

	//add services
	emailSvc := mail.New("ver2@gmail.com", "587")
	ctx := context.Background()

	fcmSvc, err := fcm.New(ctx)

	if err != nil {
		return nil, err
	}
	notify.UseServices(emailSvc, fcmSvc)

	adapters := map[string]notifier.NotifyInterface{
		"email": emailSvc,
		"fcm":   fcmSvc,
	}

	notifyService := notifier.NewNotifyService(repo, adapters)
	return notifyService, nil
}
