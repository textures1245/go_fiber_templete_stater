package usecase

import (
	"context"
	"net/http"

	"github.com/textures1245/go-template/internal/file"
	fileEntities "github.com/textures1245/go-template/internal/file/entities"
	"github.com/textures1245/go-template/internal/product"
	"github.com/textures1245/go-template/internal/product/entities"
	"github.com/textures1245/go-template/pkg/apperror"
)

type prodUse struct {
	ProductRepo product.ProductRepository
	FileRepo    file.FileRepository
}

func NewProductUsecase(prodRepo product.ProductRepository, fileRepo file.FileRepository) product.ProductUsecase {
	return &prodUse{
		ProductRepo: prodRepo,
		FileRepo:    fileRepo,
	}
}

func (u *prodUse) OnCreateProduct(ctx context.Context, prodReq *entities.ProductCreated, fileReq *fileEntities.FileUploaderReq) (int, *apperror.CErr) {
	// Upload the image first

	if fileReq != nil {
		fileId, err := u.FileRepo.CreateFile(ctx, fileReq)

		if err != nil {
			status, cErr := apperror.CustomSqlExecuteHandler("File", err)
			return status, cErr
		}

		prodReq.FileId = fileId
	}

	// Create the product
	if _, err := u.ProductRepo.CreateProduct(ctx, prodReq); err != nil {
		status, cErr := apperror.CustomSqlExecuteHandler("Product", err)
		return status, cErr
	}

	return http.StatusOK, nil
}
