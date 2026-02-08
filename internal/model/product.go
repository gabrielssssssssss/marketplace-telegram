package model

import (
	"encoding/json"
	"time"
)

type Product struct {
	ID        string          `json:"product_id"`
	Price     int64           `json:"price"`
	Details   json.RawMessage `json:"details"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type ProductResponse struct {
	Message string  `json:"message"`
	Data    Product `json:"data"`
}
