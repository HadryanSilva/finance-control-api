package model

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	gorm.Model
	Name    string
	Balance float64
}

type AccountResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	Name      string    `json:"name"`
	Balance   float64   `json:"balance"`
}
