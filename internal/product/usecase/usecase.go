package usecase

import (
	"context"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/textures1245/go-template/internal/file"
	fileEntities "github.com/textures1245/go-template/internal/file/entities"
	"github.com/textures1245/go-template/internal/product"
	"github.com/textures1245/go-template/internal/product/dtos"
	"github.com/textures1245/go-template/internal/product/entities"
	"github.com/textures1245/go-template/pkg/apperror"
	"github.com/textures1245/go-template/pkg/utils"
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
	log.Info("Creating product", prodReq)
	if _, err := u.ProductRepo.CreateProduct(ctx, prodReq); err != nil {
		status, cErr := apperror.CustomSqlExecuteHandler("Product", err)
		return status, cErr
	}

	return http.StatusOK, nil
}

func (u *prodUse) OnGetProducts(c *fiber.Ctx, ctx context.Context) ([]*dtos.ProductDataRes, int, *apperror.CErr) {
	products, err := u.ProductRepo.GetProducts(ctx)
	if err != nil {
		status, cErr := apperror.CustomSqlExecuteHandler("Product", err)
		return nil, status, cErr
	}
	for _, v := range products {
		if v.FileName != nil {
			f := fileEntities.File{
				FileName: *v.FileName,
				FileData: *v.FileBase64Url,
				FileType: *v.FileType,
			}

			_, fPathDat, status, errOnDecode := f.DecodeBlobToFile(c, true)
			if errOnDecode != nil {
				return nil, status, errOnDecode
			}
			v.FileUrl = fPathDat
			v.FileBase64Url = nil
		}
	}

	return products, http.StatusOK, nil
}

func (u *prodUse) ExportDataAsExcel(c *fiber.Ctx, ctx context.Context) (*dtos.ProductToExcelRes, int, *apperror.CErr) {

	products, err := u.ProductRepo.GetProducts(ctx)
	if err != nil {
		status, cErr := apperror.CustomSqlExecuteHandler("Product", err)
		return nil, status, cErr
	}

	if len(products) == 0 {
		return nil, http.StatusNotFound, apperror.NewCErr(
			errors.New("Product not found"),
			nil,
		)
	}

	for _, v := range products {
		if v.FileName != nil {
			f := fileEntities.File{
				FileName: *v.FileName,
				FileData: *v.FileBase64Url,
				FileType: *v.FileType,
			}

			_, fPathDat, status, errOnDecode := f.DecodeBlobToFile(c, true)
			if errOnDecode != nil {
				return nil, status, errOnDecode
			}
			v.FileUrl = fPathDat
			v.FileBase64Url = nil
		}
	}

	excelDat := utils.Excel[dtos.ProductDataRes]{
		Data: products,
	}

	res, errOnExport := excelDat.ExportData()
	if errOnExport != nil {
		return nil, http.StatusInternalServerError, errOnExport
	}

	return &dtos.ProductToExcelRes{
		FileName:   res.FileName,
		FileBuffer: res.FileBuffer,
	}, http.StatusOK, nil

}
