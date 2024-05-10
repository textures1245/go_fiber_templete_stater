package router

import (
	"runtime"
	"strings"

	"github.com/textures1245/go-template/controller"
	"github.com/textures1245/go-template/handler"
	"github.com/textures1245/go-template/repository"
	"github.com/textures1245/go-template/service"

	authV1 "github.com/textures1245/go-template/internal/auth/controller/http/v1"
	fileV1 "github.com/textures1245/go-template/internal/file/controller/http/v1"
	prodV1 "github.com/textures1245/go-template/internal/product/controller/http/v1"
	userV1 "github.com/textures1245/go-template/internal/user/controller/http/v1"

	// "github.com/textures1245/go-template/config"
	// "github.com/textures1245/go-template/pkg/datasource"

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

	api := app.Group("/", func(c *fiber.Ctx) error {
		if !strings.Contains(c.Request().URI().String(), "/ping") {
			log.Infof("all : %v", c.Request().URI().String())
		}
		return c.Next()
	})
	api.Get("/api/something", controller.SampleControllerFunction)
	api.Get("/api/ping", controller.Ping)

	userV1.UseUserRoute(db, app)
	authV1.UseAuthRoute(db, app)
	fileV1.UseFileRoute(db, app)
	prodV1.UseProductRoute(db, app)

	callback := app.Group("/callback", func(c *fiber.Ctx) error {
		log.Infof("callback : %v", c.Request().URI().String())
		return c.Next()
	})

	callback.Get("", controller.SampleControllerFunction)

}
