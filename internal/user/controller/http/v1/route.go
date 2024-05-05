package v1

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/textures1245/go-template/internal/user/repository"
	"github.com/textures1245/go-template/internal/user/usecase"
	"github.com/textures1245/go-template/pkg/middleware"
)

func UseUserRoute(db *sqlx.DB, app *fiber.App) {
	userR := app.Group("/api/user", func(c *fiber.Ctx) error {
		if !strings.Contains(c.Request().URI().String(), "/ping") {
			log.Infof("all : %v", c.Request().URI().String())
		}
		return c.Next()
	})

	// TODO: Do PUT,DELETE,GET,POST on /api/user

	userRepo := repository.NewUserRepository(db)
	userUse := usecase.NewUserUsecase(userRepo)
	userConn := NewUserController(userUse)

	userR.Get("/get-users", middleware.JwtAuthentication(), userConn.FetchUsers)
	userR.Get("/:user_id", middleware.JwtAuthentication(), userConn.FetchUserById)
	userR.Patch("/:user_id", middleware.JwtAuthentication(), userConn.UpdateUserById)
	userR.Delete("/:user_id", middleware.JwtAuthentication(), userConn.DeleteUserById)

}
