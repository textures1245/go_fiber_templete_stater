package v1

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/textures1245/go-template/internal/auth/repository"
	"github.com/textures1245/go-template/internal/auth/usecase"
	userRepo "github.com/textures1245/go-template/internal/user/repository"
)

func UseAuthRoute(db *sqlx.DB, app *fiber.App) {
	authR := app.Group("/api/v1/auth", func(c *fiber.Ctx) error {
		if !strings.Contains(c.Request().URI().String(), "/ping") {
			log.Infof("all : %v", c.Request().URI().String())
		}
		return c.Next()
	})

	authRepo := repository.NewAuthRepository(db)
	userRepo := userRepo.NewUserRepository(db)
	userUse := usecase.NewAuthService(authRepo, userRepo)
	authConn := NewAuthHandler(userUse)

	authR.Post("/customer/login", authConn.Login)
	authR.Post("/customer/register", authConn.Register)
}
