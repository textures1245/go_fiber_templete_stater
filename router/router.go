package router

import (
	"payso-simple-noti/controller"
	"payso-simple-noti/handler"
	"payso-simple-noti/service"
	"runtime"
	"strings"

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

	api := app.Group("/", func(c *fiber.Ctx) error {
		if !strings.Contains(c.Request().URI().String(), "/ping") {
			log.Infof("all : %v", c.Request().URI().String())
		}

		return c.Next()
	})

	api.Get("/api/something", controller.SampleControllerFunction)
	api.Get("/api/ping", controller.Ping)

	callback := app.Group("/callback", func(c *fiber.Ctx) error {
		log.Infof("callback : %v", c.Request().URI().String())
		return c.Next()
	})

	callback.Get("", controller.SampleControllerFunction)

}
