package product

import (
	"context"

	fileEntities "github.com/textures1245/go-template/internal/file/entities"
	"github.com/textures1245/go-template/internal/product/entities"
	"github.com/textures1245/go-template/pkg/apperror"
)

type ProductUsecase interface {
	OnCreateProduct(ctx context.Context, prodReq *entities.ProductCreated, fileReq *fileEntities.FileUploaderReq) (int, *apperror.CErr)
}
