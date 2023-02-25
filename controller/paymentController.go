package controller

import (
	"net/http"
	"payso/payment-service/model"
	"payso/payment-service/service"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type PaymentController interface {
	SampleControllerFunction(c *fiber.Ctx) error

	Ping(c *fiber.Ctx) error
	Payment(c *fiber.Ctx) error
}

type paymentController struct {
	paymentService service.PaymentService
}

func NewPaymentController(paymentService service.PaymentService) paymentController {
	return paymentController{paymentService}
}

func (obj paymentController) SampleControllerFunction(c *fiber.Ctx) error {
	/** Define log component **/
	log := log.WithFields(log.Fields{
		"component": "SampleController",
		"funciton":  "SampleControllerFunction",
	})
	log.Infof("-= Create Controller Logic Here =-")
	serviceResult, err := obj.paymentService.SampleServiceFunction(model.SampleModel{})
	if err != nil {
		log.Error(err)
		c.Response().SetStatusCode(http.StatusInternalServerError)
		return c.JSON(fiber.Map{"message": "ERROR"})
	}

	log.Debugf("serviceResult : %v", serviceResult)
	c.Response().SetStatusCode(http.StatusOK)
	return c.JSON(model.SampleModel{})
}

// Ping Healthcheck
func (obj paymentController) Ping(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "payment-api 0.99.0"})
}

func (obj paymentController) Payment(c *fiber.Ctx) error {
	log := log.WithFields(log.Fields{
		"component": "PaymentController",
		"funciton":  "Payment",
	})
	input := make(map[string]string)
	err := c.BodyParser(&input)
	if err != nil {
		log.Error(err)
		c.Response().SetStatusCode(http.StatusInternalServerError)
		return c.JSON(fiber.Map{"message": "mal format input"})

	}
	log.Infof("input : %#v", input)
	//Check valid input

	c.Response().SetStatusCode(http.StatusOK)
	return c.JSON(fiber.Map{"forwardurl": "https://www.paysolutions.asia", "QRBase64": "ddd", "QRData": "xxxxx"})
}
