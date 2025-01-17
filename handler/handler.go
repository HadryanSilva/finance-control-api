package handler

import (
	"github.com/HadryanSilva/finance-control-api/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)
