package v1

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/textures1245/go-template/internal/user"
	"github.com/textures1245/go-template/internal/user/entities"
)

type userCon struct {
	userUse user.UserUsecase
}

func NewUserController(userUse user.UserUsecase) *userCon {
	return &userCon{
		userUse: userUse,
	}
}

func (con *userCon) UserLogin(c *fiber.Ctx) error {
	req := new(entities.UserLogin)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(fiber.StatusBadRequest),
			"status_code": fiber.StatusBadRequest,
			"message":     err.Error(),
			"result":      nil,
		})
	}
	log.Info(req)

	res, err := con.userUse.OnFindUser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(fiber.StatusNotFound),
			"status_code": fiber.StatusNotFound,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      http.StatusText(fiber.StatusOK),
		"status_code": fiber.StatusOK,
		"message":     "",
		"result":      res,
	})
}

func (con *userCon) FetchUser(c *fiber.Ctx) error {
	users, err := con.userUse.OnFetchUser()
	log.Debug("Logged")
	if err != nil {
		log.Debug("Logged 2")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(fiber.StatusNotFound),
			"status_code": fiber.StatusBadRequest,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      http.StatusText(fiber.StatusOK),
		"status_code": fiber.StatusOK,
		"message":     "",
		"result":      users,
	})
}
