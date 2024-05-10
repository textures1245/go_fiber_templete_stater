package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/textures1245/go-template/internal/product"
	"github.com/textures1245/go-template/internal/product/dtos"
	"github.com/textures1245/go-template/internal/product/entities"
	"github.com/textures1245/go-template/internal/product/repository/repository_query"
	"github.com/textures1245/go-template/pkg/utils"
)

type productRepo struct {
	Db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) product.ProductRepository {
	return &productRepo{
		Db: db,
	}
}

func (r *productRepo) CreateProduct(ctx context.Context, product *entities.ProductCreated) (*int64, error) {

	args := utils.Array{
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
		product.FileId,
	}

	res, err := r.Db.ExecContext(ctx, repository_query.CreateProduct, args...)
	if err != nil {
		return nil, err
	}

	createID, err := res.LastInsertId()
	if err != nil {
		return nil, sql.ErrNoRows
	}

	return &createID, nil
}

func (r *productRepo) GetProducts(ctx context.Context) ([]*dtos.ProductDataRes, error) {
	var products []*dtos.ProductDataRes

	err := r.Db.SelectContext(ctx, &products, repository_query.GetProductsLeftJoinFile)
	if err != nil {
		return nil, err
	}

	return products, nil
}
