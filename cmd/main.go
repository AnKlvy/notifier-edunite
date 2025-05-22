package main

import (
	"context"
	"database/sql"
	"github.com/AnKlvy/notifier-edunite/internal/config"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	"time"
)

const version = "1.0.0"

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}
	db, err := openDB(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
	ntf, err := initNotify(db)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("database connection pool established")

	port := ":" + strconv.Itoa(cfg.Port)

	grpcServer := NewGRPCServer(port, *ntf)

	log.Println("starting server", map[string]string{
		"addr": grpcServer.addr,
		"env":  cfg.Env,
	})

	// Запускаем gRPC-сервер в отдельной горутине
	go func() {
		if serveErr := grpcServer.Run(); serveErr != nil {
			log.Fatal(serveErr)
		}
	}()

	// Ждём сигнала завершения (Ctrl+C или SIGTERM в Kubernetes)
	waitForShutdown(grpcServer.server)
	log.Println("Server gracefully stopped")
}

func openDB(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.Db.Dsn)
	if err != nil {
		return nil, err
	}

	// Устанавливаем максимальное количество открытых (используемых + свободных) соединений в пуле.
	// Если передано значение меньше или равное 0, ограничение не устанавливается.
	db.SetMaxOpenConns(cfg.Db.MaxOpenConns)

	// Устанавливаем максимальное количество свободных соединений в пуле.
	// Если передано значение меньше или равное 0, ограничение не устанавливается.
	db.SetMaxIdleConns(cfg.Db.MaxIdleConns)

	// Используем функцию time.ParseDuration() для преобразования строки с таймаутом простоя
	// в тип time.Duration.
	duration, err := time.ParseDuration(cfg.Db.MaxIdleTime)
	if err != nil {
		return nil, err
	}

	// Устанавливаем максимальное время простоя соединений.
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Проверяем соединение с базой данных.
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
