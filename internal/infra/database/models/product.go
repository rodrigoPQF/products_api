package models

import "time"

type Product struct {
	Id string `db:"id,omitempty"`
	Name string `db:"name"`
	Price float64 `db:"price"`
	Description string `db:"description"`
	CreatedAt time.Time `db:"created_at,omitempty"`
	UpdatedAt time.Time `db:"updated_at,omitempty"`
}