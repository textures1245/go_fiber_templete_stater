package file

import (
	"context"

	"github.com/textures1245/go-template/internal/file/dtos"
	"github.com/textures1245/go-template/internal/file/entities"
	"github.com/textures1245/go-template/pkg/apperror"
)

type FileUsecase interface {
	OnUploadFile(ctx context.Context, req *entities.FileUploaderReq) (int, *apperror.CErr)
	GetSourceFiles(ctx context.Context) ([]*dtos.FileSourceDataRes, int, *apperror.CErr)
}
