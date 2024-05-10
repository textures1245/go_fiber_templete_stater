package usecase

import (
	"context"

	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/textures1245/go-template/internal/file"
	"github.com/textures1245/go-template/internal/file/dtos"
	"github.com/textures1245/go-template/internal/file/entities"
	"github.com/textures1245/go-template/pkg/apperror"
)

type fileUsecase struct {
	fileRepo file.FileRepository
}

func NewFileUsecase(fileRepo file.FileRepository) file.FileUsecase {
	return &fileUsecase{
		fileRepo: fileRepo,
	}
}

func (u *fileUsecase) OnUploadFile(c *fiber.Ctx, ctx context.Context, req *entities.FileUploaderReq) (*dtos.FileSourceDataRes, int, *apperror.CErr) {
	fileId, errCreateFile := u.fileRepo.CreateFile(ctx, req)

	if errCreateFile != nil {
		status, cErr := apperror.CustomSqlExecuteHandler("File", errCreateFile)
		return nil, status, cErr
	}

	file, err := u.fileRepo.GetFileById(ctx, fileId)
	if err != nil {
		status, cErr := apperror.CustomSqlExecuteHandler("File", err)
		return nil, status, cErr
	}

	_, fPathDat, status, errOnDecode := file.DecodeBlobToFile(c, true)
	if errOnDecode != nil {
		return nil, status, errOnDecode
	}

	filesRes := &dtos.FileSourceDataRes{
		FileName: file.FileName,
		// FileBase64URL: base64urlRes,
		FilePathData: *fPathDat,
		FileType:     file.FileType,
		CreatedAt:    file.CreatedAt,
		UpdatedAt:    file.UpdatedAt,
	}

	return filesRes, http.StatusOK, nil
}

func (u *fileUsecase) GetSourceFiles(c *fiber.Ctx, ctx context.Context) ([]*dtos.FileSourceDataRes, int, *apperror.CErr) {
	files, err := u.fileRepo.GetFiles(ctx)
	if err != nil {
		status, cErr := apperror.CustomSqlExecuteHandler("File", err)
		return nil, status, cErr
	}

	filesRes := []*dtos.FileSourceDataRes{}
	for _, file := range files {
		_, fPathDat, status, errOnDecode := file.DecodeBlobToFile(c, true)
		if errOnDecode != nil {
			return nil, status, errOnDecode
		}

		filesRes = append(filesRes, &dtos.FileSourceDataRes{
			FileName: file.FileName,
			// FileBase64URL: base64urlRes,
			FilePathData: *fPathDat,
			FileType:     file.FileType,
			CreatedAt:    file.CreatedAt,
			UpdatedAt:    file.UpdatedAt,
		})
	}

	return filesRes, http.StatusOK, nil
}
