package repository

import (
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)

	db, err := sqlx.Open("pgx", connStr)
	if err != nil {
		logrus.Errorf("NewPostgresDB: failed to open connection: %s", err.Error())
		return nil, err
	}

	if err = db.Ping(); err != nil {
		logrus.Errorf("NewPostgresDB: failed to ping database: %s", err.Error())
		return nil, err
	}

	logrus.Info("NewPostgresDB: database connection established")
	return db, nil
}
