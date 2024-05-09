package product

import (
	"context"

	"github.com/textures1245/go-template/internal/product/entities"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, req *entities.ProductCreated) (*int64, error)
}
