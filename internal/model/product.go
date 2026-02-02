package model

import "time"

type Product struct {
	ID        string    `json:"product_id"`
	Details   any       `json:"details"`
	Price     int64     `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
