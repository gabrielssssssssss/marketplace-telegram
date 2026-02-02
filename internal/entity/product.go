package entity

import "time"

type Product struct {
	ID        string
	Details   any
	Price     int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
