package model

import "time"

type Cart struct {
	ID        string    `json:"cart_id"`
	UserID    int64     `json:"user_id"`
	ProductID string    `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
