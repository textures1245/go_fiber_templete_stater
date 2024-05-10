package v1

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	fileRepo "github.com/textures1245/go-template/internal/file/repository"
	"github.com/textures1245/go-template/internal/product/repository"
	"github.com/textures1245/go-template/internal/product/usecase"
	// "github.com/textures1245/go-template/pkg/middleware"
)

func UseProductRoute(db *sqlx.DB, app *fiber.App) {
	authR := app.Group("/api/v1/product", func(c *fiber.Ctx) error {
		if !strings.Contains(c.Request().URI().String(), "/ping") {
			log.Infof("all : %v", c.Request().URI().String())
		}
		return c.Next()
	})

	fileRepo := fileRepo.NewFileRepository(db)
	prodRepo := repository.NewProductRepository(db)
	prodUse := usecase.NewProductUsecase(prodRepo, fileRepo)
	prodConn := NewProductHandler(prodUse)

	authR.Post("/add-products", prodConn.CreateProducts)
	authR.Get("/get-products", prodConn.GetProducts)
	authR.Get("/export-to-excel", prodConn.ExportDataAsExcel)
}
