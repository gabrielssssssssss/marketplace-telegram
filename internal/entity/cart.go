package entity

import "time"

type Cart struct {
	ID        string
	UserID    int64
	ProductID string
	CreatedAt time.Time
	UpdatedAt time.Time
}
