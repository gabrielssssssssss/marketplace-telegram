package entity

import (
	"encoding/json"
	"time"
)

type Product struct {
	ID             string
	PrivateDetails *json.RawMessage
	PublicDetails  *json.RawMessage
	Price          *int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
