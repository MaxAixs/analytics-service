package database

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
)

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg DBConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))
	if err != nil {
		logrus.Printf("error connecting to database: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logrus.Printf("error pinging database: %v", err)
	}

	return db, nil
}
