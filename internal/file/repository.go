package file

import (
	"context"

	"github.com/textures1245/go-template/internal/file/entities"
)

type FileRepository interface {
	CreateFile(ctx context.Context, file *entities.FileUploaderReq) error
	GetFiles(ctx context.Context) ([]*entities.File, error)
}
