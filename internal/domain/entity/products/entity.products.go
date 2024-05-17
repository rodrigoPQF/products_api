package products

import "time"

type ProductEntity struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}