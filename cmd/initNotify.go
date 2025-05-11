package main

import (
	"database/sql"
	"github.com/AnKlvy/notifier-edunite/internal/database"
	"github.com/AnKlvy/notifier-edunite/internal/services/notifier"
	"github.com/AnKlvy/notifier-edunite/internal/services/notifier/email"
	"github.com/AnKlvy/notifier-edunite/internal/services/notifier/firebase"
)

func initNotify(db *sql.DB) (*notifier.NotifyService, error) {
	//init repo interface
	repo := database.NewNotifier(db)

	//add services
	emailSvc := email.InitEmail()

	firebaseSvc, err := firebase.InitFirebase()
	if err != nil {
		return nil, err
	}

	adapters := map[string]notifier.NotifyInterface{
		"email":    emailSvc,
		"firebase": firebaseSvc,
	}

	notifyService := notifier.NewNotifyService(repo, adapters)
	return notifyService, nil
}
