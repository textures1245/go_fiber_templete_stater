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

	api := app.Group("/api", func(c *fiber.Ctx) error {
		if !strings.Contains(c.Request().URI().String(), "/ping") {
			log.Info(c.Request().URI().String())
		}

		return c.Next()
	})

	api.Get("/something", controller.SampleControllerFunction)
	api.Get("/ping", controller.Ping)

}
