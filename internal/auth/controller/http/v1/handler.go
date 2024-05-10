package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/textures1245/go-template/internal/auth"
	"github.com/textures1245/go-template/internal/auth/entities"
	_userEntities "github.com/textures1245/go-template/internal/user/entities"
	"github.com/textures1245/go-template/pkg/utils"
)

type authConn struct {
	AuthUse auth.AuthUsecase
}

func NewAuthHandler(authUse auth.AuthUsecase) *authConn {
	return &authConn{
		AuthUse: authUse,
	}
}

func (a *authConn) Login(c *fiber.Ctx) error {
	req := new(entities.UsersCredentials)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid request body",
			"raw_message": err.Error(),
			"result":      nil,
		})
	}

	errOnValidate := utils.SchemaValidator(req)
	if errOnValidate != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid validated on schema body",
			"raw_message": errOnValidate.RawError.Error(),
			"result":      nil,
		})
	}

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	res, status, err := a.AuthUse.Login(ctx, req)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      http.StatusText(status),
			"status_code": status,
			"message":     err.CError.Error(),
			"raw_message": err.RawError.Error(),
			"result":      nil,
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"status":      http.StatusText(status),
		"status_code": status,
		"message":     "",
		"result":      res,
	})
}

func (a *authConn) Register(c *fiber.Ctx) error {
	req := new(_userEntities.UserCreatedReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid request body",
			"raw_message": err.Error(),
			"result":      nil,
		})
	}

	errOnValidate := utils.SchemaValidator(req)
	if errOnValidate != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid validated on schema body",
			"raw_message": errOnValidate.RawError.Error(),
			"result":      nil,
		})
	}

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	res, status, err := a.AuthUse.Register(ctx, req)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      http.StatusText(status),
			"status_code": status,
			"message":     err.CError.Error(),
			"raw_message": err.RawError.Error(),
			"result":      nil,
		})
	}

	// TODO: Created Recovery when failed to send email
	email := &utils.Email{
		From:    "comcamp.22nd@gmail.com",
		To:      "sirprak1245@gmail.com",
		Subject: "Register Successfully",
		Body:    "Welcome to our platform",
	}
	if errOnSend := email.Send(); errOnSend != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusInternalServerError),
			"status_code": http.StatusInternalServerError,
			"message":     "error, failed to send email",
			"raw_message": errOnSend.Error(),
			"result":      nil,
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"status":      http.StatusText(status),
		"status_code": status,
		"message":     "",
		"result":      res,
	})
}
