package entity

import (
	"time"
)

type Users struct {
	UserId      string
	DisplayName string
	Username    string
	Balance     float64
	RecoveryKey string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
