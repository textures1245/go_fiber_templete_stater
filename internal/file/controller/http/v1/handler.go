package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/textures1245/go-template/internal/file"
	"github.com/textures1245/go-template/internal/file/entities"
	"github.com/textures1245/go-template/pkg/utils"
	// "github.com/textures1245/go-template/internal/file/usecase"
)

type fileConn struct {
	FileUse file.FileUsecase
}

func NewFileHandler(authUse file.FileUsecase) *fileConn {
	return &fileConn{
		FileUse: authUse,
	}
}

func (h *fileConn) UploadFile(c *fiber.Ctx) error {
	req := new(entities.FileUploaderReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid request body",
			"raw_message": err.Error(),
			"result":      nil,
		})
	}

	errOnValidate := utils.SchemaValidator(req)
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

	res, status, err := h.FileUse.OnUploadFile(c, ctx, req)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      http.StatusText(status),
			"status_code": status,
			"message":     err.CError.Error(),
			"raw_message": err.RawError.Error(),
			"result":      nil,
		})
	}

	if res.FileType == "PDF" {
		return c.SendFile(res.FilePathData)
	}
	return c.Status(status).JSON(fiber.Map{
		"status":      http.StatusText(status),
		"status_code": status,
		"message":     "",
		"result":      res,
	})

}

func (h *fileConn) GetSourceFiles(c *fiber.Ctx) error {
	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	files, status, err := h.FileUse.GetSourceFiles(c, ctx)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      http.StatusText(status),
			"status_code": status,
			"message":     err.CError.Error(),
			"raw_message": err.RawError.Error(),
			"result":      nil,
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"status":      http.StatusText(status),
		"status_code": status,
		"message":     "",
		"result":      files,
	})
}
