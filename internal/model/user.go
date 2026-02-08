package model

import (
	"time"
)

type User struct {
	UserId      int64     `json:"user_id"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	Username    string    `json:"username"`
	Role        string    `json:"role"`
	Balance     float64   `json:"balance"`
	RecoveryKey string    `json:"recovery_key"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserResponse struct {
	Message string `json:"message"`
	Data    User   `json:"data"`
}
