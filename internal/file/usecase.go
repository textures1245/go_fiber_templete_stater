package file

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/textures1245/go-template/internal/file/dtos"
	"github.com/textures1245/go-template/internal/file/entities"
	"github.com/textures1245/go-template/pkg/apperror"
)

type FileUsecase interface {
	OnUploadFile(c *fiber.Ctx, ctx context.Context, req *entities.FileUploaderReq) (*dtos.FileSourceDataRes, int, *apperror.CErr)
	GetSourceFiles(c *fiber.Ctx, ctx context.Context) ([]*dtos.FileSourceDataRes, int, *apperror.CErr)
}
