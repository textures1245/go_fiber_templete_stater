package entities

import "github.com/textures1245/go-template/internal/file/entities"

type ProductCreatedReq struct {
	Name        string                    `json:"product_name" db:"product_name" form:"product_name" binding:"required"`
	Description string                    `json:"product_description" db:"product_description"  form:"product_description" binding:"required"`
	Price       float64                   `json:"product_price"  db:"product_price" form:"product_price" binding:"required"`
	Stock       int64                     `json:"stock_qty" db:"stock_qty"  form:"stock_qty" binding:"required"`
	File        *entities.FileUploaderReq `json:"file" form:"file" `
}

// create Product request only
type ProductCreated struct {
	Name        string  `json:"product_name" form:"product_name" db:"product_name" binding:"required"`
	Description string  `json:"product_description" form:"product_description" db:"product_description" binding:"required"`
	Price       float64 `json:"product_price" form:"product_price" db:"product_price" binding:"required"`
	Stock       int64   `json:"stock_qty" form:"stock_qty" db:"stock_qty" binding:"required"`
	FileId      *int64  `json:"file_id" db:"file_id"`
}
