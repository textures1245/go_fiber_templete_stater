package product

import (
	"context"

	"github.com/gofiber/fiber/v2"
	fileEntities "github.com/textures1245/go-template/internal/file/entities"
	"github.com/textures1245/go-template/internal/product/dtos"
	"github.com/textures1245/go-template/internal/product/entities"
	"github.com/textures1245/go-template/pkg/apperror"
)

type ProductUsecase interface {
	OnCreateProduct(ctx context.Context, prodReq *entities.ProductCreated, fileReq *fileEntities.FileUploaderReq) (int, *apperror.CErr)
	OnGetProducts(c *fiber.Ctx, ctx context.Context) ([]*dtos.ProductDataRes, int, *apperror.CErr)
	ExportDataAsExcel(c *fiber.Ctx, ctx context.Context) (*dtos.ProductToExcelRes, int, *apperror.CErr)
}
