package config

import (
	"fmt"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	db, err = InitPostgres()
	if err != nil {
		return fmt.Errorf("failed to initialize Postgres: %v", err)
	}

	return nil
}

func GetPostgres() *gorm.DB {
	return db
}

func GetLogger(s string) *Logger {
	logger = newLogger(s)
	return logger
}
