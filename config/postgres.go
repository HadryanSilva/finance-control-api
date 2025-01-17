package config

import (
	"github.com/HadryanSilva/finance-control-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres= password=postgres dbname=postgres search_path=finance_db"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.Account{})
	if err != nil {
		return nil, err
	}

	return db, err
}
