package router

import (
	"runtime"
	"strings"

	"github.com/textures1245/go-template/controller"
	"github.com/textures1245/go-template/handler"
	"github.com/textures1245/go-template/repository"
	"github.com/textures1245/go-template/service"

	_userConn "github.com/textures1245/go-template/internal/user/controller/http/v1"
	_userRepo "github.com/textures1245/go-template/internal/user/repository"
	_userUsecase "github.com/textures1245/go-template/internal/user/usecase"

	_authConn "github.com/textures1245/go-template/internal/auth/controller/http/v1"
	_authRepo "github.com/textures1245/go-template/internal/auth/repository"
	_authUsecase "github.com/textures1245/go-template/internal/auth/usecase"

	// "github.com/textures1245/go-template/config"
	// "github.com/textures1245/go-template/pkg/datasource"
	"github.com/textures1245/go-template/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func SetupRoutes(app *fiber.App) {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})

	controller := controller.NewSampleController(service.NewSampleService(handler.NewSimpleHandler()))

	// set user conn
	// db, err := datasource.NewDB(config.LoadDBconfig())
	// if err != nil {
	// 	log.Errorf("Error : %v", err)
	// }
	// log.Info(db)

	db := repository.GetDb()

	userRepo := _userRepo.NewUserRepository(db)
	userService := _userUsecase.NewUserUsecase(userRepo)
	userConn := _userConn.NewUserController(userService)

	api := app.Group("/", func(c *fiber.Ctx) error {
		if !strings.Contains(c.Request().URI().String(), "/ping") {
			log.Infof("all : %v", c.Request().URI().String())
		}
		return c.Next()
	})

	userR := app.Group("/api/user", func(c *fiber.Ctx) error {
		if !strings.Contains(c.Request().URI().String(), "/ping") {
			log.Infof("all : %v", c.Request().URI().String())
		}
		return c.Next()
	})

	authR := app.Group("/api/auth", func(c *fiber.Ctx) error {
		if !strings.Contains(c.Request().URI().String(), "/ping") {
			log.Infof("all : %v", c.Request().URI().String())
		}
		return c.Next()
	})

	authRepo := _authRepo.NewAuthRepository(db)
	authService := _authUsecase.NewAuthService(authRepo, userRepo)
	authConn := _authConn.NewAuthHandler(authService)

	api.Get("/api/something", controller.SampleControllerFunction)
	api.Get("/api/ping", controller.Ping)

	userR.Get("/get-users", middleware.JwtAuthentication(), userConn.FetchUsers)

	authR.Post("/customer/login", authConn.Login)
	authR.Post("/customer/register", authConn.Register)

	callback := app.Group("/callback", func(c *fiber.Ctx) error {
		log.Infof("callback : %v", c.Request().URI().String())
		return c.Next()
	})

	callback.Get("", controller.SampleControllerFunction)

}
