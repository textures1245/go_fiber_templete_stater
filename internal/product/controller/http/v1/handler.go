package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	fileEntities "github.com/textures1245/go-template/internal/file/entities"
	"github.com/textures1245/go-template/internal/product"
	"github.com/textures1245/go-template/internal/product/entities"
	"github.com/textures1245/go-template/pkg/utils"
)

type prodConn struct {
	prodUse product.ProductUsecase
}

func NewProductHandler(prodUse product.ProductUsecase) *prodConn {
	return &prodConn{
		prodUse: prodUse,
	}
}

func (h *prodConn) CreateProducts(c *fiber.Ctx) error {
	req := make([]*entities.ProductCreatedReq, 0)
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid request body",
			"raw_message": err.Error(),
			"result":      nil,
		})
	}

	errOnValidate := utils.SchemaValidator(req[0])
	if errOnValidate != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid validated on schema body",
			"raw_message": errOnValidate.RawError.Error(),
			"result":      nil,
		})
	}

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	for _, v := range req {
		prodReq := &entities.ProductCreated{
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
			Stock:       v.Stock,
			FileId:      nil,
		}

		var fileReq *fileEntities.FileUploaderReq
		if v.File != nil {
			fileReq = &fileEntities.FileUploaderReq{
				FileName: v.File.FileName,
				FileType: v.File.FileType,
				FileData: v.File.FileData,
			}
		}

		status, err := h.prodUse.OnCreateProduct(ctx, prodReq, fileReq)
		if err != nil {
			return c.Status(status).JSON(fiber.Map{
				"status":      http.StatusText(status),
				"status_code": status,
				"message":     err.CError.Error(),
				"raw_message": err.RawError.Error(),
				"result":      nil,
			})
		}
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusText(http.StatusOK),
		"status_code": http.StatusOK,
		"message":     "",
		"result":      "",
	})

}
