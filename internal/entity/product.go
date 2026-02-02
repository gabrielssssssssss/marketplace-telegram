package entity

import "time"

type Product struct {
	ID        *int64
	Details   *map[string]any
	Price     *int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
