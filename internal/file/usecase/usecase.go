package usecase

import (
	"context"
	"errors"
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

	var (
		// base64urlRes string
		fPathDatRes string
	)
	switch file.FileType {
	case "PNG":
		_, fPathDat, err := file.Base64toPng(c)
		if err != nil {
			status, cErr := apperror.CustomSqlExecuteHandler("File", err)
			return nil, status, cErr
		}
		// base64urlRes = *base64url
		fPathDatRes = *fPathDat
	case "JPG":
		_, fPathDat, err := file.Base64toJpg(c)
		if err != nil {
			status, cErr := apperror.CustomSqlExecuteHandler("File", err)
			return nil, status, cErr
		}
		// base64urlRes = *base64url
		fPathDatRes = *fPathDat
	case "PDF":
		_, fPathDat, err := file.Base64toFile(c, false)
		if err != nil {
			status, cErr := apperror.CustomSqlExecuteHandler("File", err)
			return nil, status, cErr
		}
		// base64urlRes = *base64url
		fPathDatRes = *fPathDat
	default:
		return nil, http.StatusBadRequest, apperror.NewCErr(errors.New("Only except for PNG and JPG for now"), errors.ErrUnsupported)
	}

	filesRes := &dtos.FileSourceDataRes{
		FileName: file.FileName,
		// FileBase64URL: base64urlRes,
		FilePathData: fPathDatRes,
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
		var (
			// base64urlRes string
			fPathDatRes string
		)
		switch file.FileType {
		case "PNG":
			_, fPathDat, err := file.Base64toPng(c)
			if err != nil {
				status, cErr := apperror.CustomSqlExecuteHandler("File", err)
				return nil, status, cErr
			}
			// base64urlRes = *base64url
			fPathDatRes = *fPathDat
		case "JPG":
			_, fPathDat, err := file.Base64toJpg(c)
			if err != nil {
				status, cErr := apperror.CustomSqlExecuteHandler("File", err)
				return nil, status, cErr
			}
			// base64urlRes = *base64url
			fPathDatRes = *fPathDat
		case "PDF":
			_, fPathDat, err := file.Base64toFile(c, true)
			if err != nil {
				status, cErr := apperror.CustomSqlExecuteHandler("File", err)
				return nil, status, cErr
			}
			// base64urlRes = *base64url
			fPathDatRes = *fPathDat
		default:
			return nil, http.StatusBadRequest, apperror.NewCErr(errors.New("Only except for PNG and JPG for now"), errors.ErrUnsupported)
		}

		filesRes = append(filesRes, &dtos.FileSourceDataRes{
			FileName: file.FileName,
			// FileBase64URL: base64urlRes,
			FilePathData: fPathDatRes,
			FileType:     file.FileType,
			CreatedAt:    file.CreatedAt,
			UpdatedAt:    file.UpdatedAt,
		})
	}

	return filesRes, http.StatusOK, nil
}
