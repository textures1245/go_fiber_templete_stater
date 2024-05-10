package repository_query

import (
	_ "embed"
)

var (
	//go:embed product/create_product.sql
	CreateProduct string

	//go:embed product/get_products_left_join_file.sql
	GetProductsLeftJoinFile string
)
