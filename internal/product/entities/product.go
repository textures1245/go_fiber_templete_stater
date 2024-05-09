package entities

import "github.com/textures1245/go-template/internal/file/entities"

type Product struct {
	Id          int64   `json:"id" db:"id"`
	Name        string  `json:"name" db:"product_name"`
	Description string  `json:"description" db:"product_description"`
	Price       float64 `json:"price" db:"product_price"`
	Stock       int64   `json:"stock" db:"stock_qty"`
	CreatedAt   string  `json:"created_at" db:"created_at"`
	UpdatedAt   string  `json:"updated_at" db:"updated_at"`

	// Image
	entities.File
}
