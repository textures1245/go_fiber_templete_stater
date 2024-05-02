package controller

import (
	"net/http"
	"runtime"
	"strings"

	"github.com/textures1245/go-template/model"
	"github.com/textures1245/go-template/service"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type SampleController interface {
	SampleControllerFunction(c *fiber.Ctx) error

	Ping(c *fiber.Ctx) error
}

type sampleController struct {
	sampleService service.SampleService
}

func NewSampleController(sampleService service.SampleService) sampleController {
	return sampleController{sampleService}
}

// Ping Healthcheck
func (obj sampleController) Ping(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "sample-api 0.99.0"})
}

func (obj sampleController) SampleControllerFunction(c *fiber.Ctx) error {
	/** Define log component **/
	_, file, _, _ := runtime.Caller(0)
	pc, _, _, _ := runtime.Caller(0)
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]

	log := log.WithFields(log.Fields{
		"component": strings.Split(file, "/")[len(strings.Split(file, "/"))-1],
		"funciton":  functionName,
	})
	log.Infof("-= Create Controller Logic Here =-")
	serviceResult, err := obj.sampleService.SampleServiceFunction(model.SampleModel{})
	if err != nil {
		log.Error(err)
		c.Response().SetStatusCode(http.StatusInternalServerError)
		return c.JSON(fiber.Map{"message": "ERROR"})
	}

	log.Debugf("serviceResult : %v", serviceResult)
	c.Response().SetStatusCode(http.StatusOK)
	return c.JSON(model.SampleModel{})
}
