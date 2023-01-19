package controller

import (
	"net/http"
	"payso/go_template/model"
	"payso/go_template/service"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type SampleController interface {
	Ping(c *fiber.Ctx) error
	SampleControllerFunction(c *fiber.Ctx) error
}

type sampleController struct {
	someService service.SampleService
}

func NewSampleController(someService service.SampleService) sampleController {
	return sampleController{someService}
}

// Ping Healthcheck
func (obj sampleController) Ping(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "sample-api 0.99.0"})
}

func (obj sampleController) SampleControllerFunction(c *fiber.Ctx) error {
	/** Define log component **/
	log := log.WithFields(log.Fields{
		"component": "SampleController",
		"funciton":  "SampleControllerFunction",
	})
	log.Infof("-= Create Controller Logic Here =-")
	serviceResult, err := obj.someService.SampleServiceFunction(model.SampleModel{})
	if err != nil {
		log.Error(err)
		c.Response().SetStatusCode(http.StatusInternalServerError)
		return c.JSON(fiber.Map{"message": "ERROR"})
	}

	log.Debugf("serviceResult : %v", serviceResult)
	c.Response().SetStatusCode(http.StatusOK)
	return c.JSON(model.SampleModel{})
}
