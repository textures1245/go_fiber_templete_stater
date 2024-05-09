package repository_query

import (
	_ "embed"
)

var (
	//go:embed product/create_product.sql
	CreateProduct string
)
