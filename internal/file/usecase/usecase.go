package usecase

import (
	"context"
	"errors"
	"net/http"

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

func (u *fileUsecase) OnUploadFile(ctx context.Context, req *entities.FileUploaderReq) (int, *apperror.CErr) {
	err := u.fileRepo.CreateFile(ctx, req)
	if err != nil {
		status, cErr := apperror.CustomSqlExecuteHandler("File", err)
		return status, cErr
	}

	return http.StatusOK, nil
}

func (u *fileUsecase) GetSourceFiles(ctx context.Context) ([]*dtos.FileSourceDataRes, int, *apperror.CErr) {
	files, err := u.fileRepo.GetFiles(ctx)
	if err != nil {
		status, cErr := apperror.CustomSqlExecuteHandler("File", err)
		return nil, status, cErr
	}

	filesRes := []*dtos.FileSourceDataRes{}
	for _, file := range files {
		var srcFile string
		switch file.FileType {
		case "PNG":
			src, err := file.Base64toPng()
			if err != nil {
				status, cErr := apperror.CustomSqlExecuteHandler("File", err)
				return nil, status, cErr
			}
			srcFile = src
		case "JPG":
			src, err := file.Base64toJpg()
			if err != nil {
				status, cErr := apperror.CustomSqlExecuteHandler("File", err)
				return nil, status, cErr
			}
			srcFile = src
		case "PDF":
			src, err := file.Base64toFile()
			if err != nil {
				status, cErr := apperror.CustomSqlExecuteHandler("File", err)
				return nil, status, cErr
			}
			srcFile = src
		default:
			return nil, http.StatusBadRequest, apperror.NewCErr(errors.New("Only except for PNG and JPG for now"), errors.ErrUnsupported)
		}

		filesRes = append(filesRes, &dtos.FileSourceDataRes{
			FileName:   file.FileName,
			FileSrcURL: srcFile,
			FileType:   file.FileType,
			CreatedAt:  file.CreatedAt,
			UpdatedAt:  file.UpdatedAt,
		})
	}

	return filesRes, http.StatusOK, nil
}
