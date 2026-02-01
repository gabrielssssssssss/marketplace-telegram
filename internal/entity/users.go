package entity

import (
	"time"
)

type Users struct {
	UserId      int64
	Firstname   string
	Lastname    string
	Username    string
	Role        string
	Balance     float64
	RecoveryKey string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
