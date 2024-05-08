package repository

import (
	"context"
	"encoding/base64"

	"github.com/jmoiron/sqlx"
	"github.com/textures1245/go-template/internal/file"
	"github.com/textures1245/go-template/internal/file/entities"
	"github.com/textures1245/go-template/internal/file/repository/repository_query"
	"github.com/textures1245/go-template/pkg/utils"

	"github.com/gofiber/fiber/v2/log"
)

type fileRepo struct {
	db *sqlx.DB
}

func NewFileRepository(db *sqlx.DB) file.FileRepository {
	return &fileRepo{
		db: db,
	}
}

func (r *fileRepo) GetFiles(ctx context.Context) ([]*entities.File, error) {
	var files []*entities.File

	err := r.db.SelectContext(ctx, &files, repository_query.GetFiles)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return files, nil
}

func (r *fileRepo) CreateFile(ctx context.Context, file *entities.FileUploaderReq) error {
	b64, err := base64.StdEncoding.DecodeString(file.FileData)
	if err != nil {
		log.Info(err)
		return err
	}

	args := utils.Array{
		file.FileName,
		b64,
		file.FileType,
	}

	res, err := r.db.ExecContext(ctx, repository_query.CreateFile, args...)
	if err != nil {
		log.Info(err)
		return err
	}
	if _, err := res.LastInsertId(); err != nil {
		log.Info(err)
		return err
	}

	return nil
}
