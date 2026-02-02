package model

import (
	"encoding/json"
	"time"
)

type Product struct {
	ID        string          `json:"product_id"`
	Details   json.RawMessage `json:"details"`
	Price     int64           `json:"price"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
