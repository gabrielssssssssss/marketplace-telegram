package model

import "time"

type Cart struct {
	ID        string    `json:"cart_id"`
	UserID    int64     `json:"user_id"`
	ProductID string    `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CartResponse struct {
	Message string `json:"message"`
	Data    Cart   `json:"data"`
}

type CartsResponse struct {
	Message string `json:"message"`
	Data    []Cart `json:"data"`
}
