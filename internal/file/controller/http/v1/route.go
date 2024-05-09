package v1

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/textures1245/go-template/internal/file/repository"
	"github.com/textures1245/go-template/internal/file/usecase"
)

func UseFileRoute(db *sqlx.DB, app *fiber.App) {
	authR := app.Group("/api/v1/file", func(c *fiber.Ctx) error {
		if !strings.Contains(c.Request().URI().String(), "/ping") {
			log.Infof("all : %v", c.Request().URI().String())
		}
		return c.Next()
	})

	fileRepo := repository.NewFileRepository(db)
	fileUse := usecase.NewFileUsecase(fileRepo)
	fileConn := NewFileHandler(fileUse)

	authR.Post("/upload", fileConn.UploadFile)
	authR.Get("/get-files", fileConn.GetSourceFiles)
}
