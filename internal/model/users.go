package model

import (
	"time"
)

type Users struct {
	UserId      int64      `json:"user_id"`
	Firstname   *string    `json:"firstname"`
	Lastname    *string    `json:"lastname"`
	Username    *string    `json:"username"`
	Balance     *float64   `json:"balance"`
	RecoveryKey *string    `json:"recovery_key"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
