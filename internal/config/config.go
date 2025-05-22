package config

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Добавляем поля maxOpenConns, maxIdleConns и maxIdleTime для хранения
// параметров конфигурации пула подключений.
type Config struct {
	Port int
	Env  string
	Db   struct {
		Dsn          string
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime  string
	}
	// Добавляем новую структуру limiter, содержащую поля для количества запросов в секунду,
	// максимального числа запросов в очереди (burst) и булево поле, которое можно использовать
	// для включения/отключения ограничения запросов.
	limiter struct {
		rps     float64
		burst   int
		enabled bool
	}
}

func LoadConfig() (*Config, error) {
	var cfg Config

	// Загружаем .env файл, но не возвращаем ошибку, если файл не найден
	// Это позволит использовать переменные окружения, даже если .env отсутствует
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		// Логируем ошибку, но продолжаем выполнение
		log.Printf("Ошибка в .env файле: %v", err)
	}

	flag.IntVar(&cfg.Port, "port", 9100, "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.Db.Dsn, "db-dsn", os.Getenv("NOTIFY_SERVICE_DB_DSN"), "PostgreSQL DSN")
	flag.IntVar(&cfg.Db.MaxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.Db.MaxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.Db.MaxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	flag.Float64Var(&cfg.limiter.rps, "limiter-rps", 2, "Rate limiter maximum requests per second")
	flag.IntVar(&cfg.limiter.burst, "limiter-burst", 4, "Rate limiter maximum burst")
	flag.BoolVar(&cfg.limiter.enabled, "limiter-enabled", true, "Enable rate limiter")

	flag.Parse()
	return &cfg, nil
}
