package model

import (
	"time"
)

type Users struct {
	UserId      string    `json:"user_id"`
	DisplayName string    `json:"display_name"`
	Username    string    `json:"username"`
	Balance     float64   `json:"balance"`
	RecoveryKey string    `json:"recovery_key"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
