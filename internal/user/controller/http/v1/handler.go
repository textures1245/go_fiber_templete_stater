package v1

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/textures1245/go-template/internal/user"
)

type userCon struct {
	userUse user.UserUsecase
}

func NewUserController(userUse user.UserUsecase) *userCon {
	return &userCon{
		userUse: userUse,
	}
}

// func (con *userCon) UserLogin(c *fiber.Ctx) error {
// 	req := new(entities.UserLoginReq)
// 	if err := c.BodyParser(req); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"status":      http.StatusText(fiber.StatusBadRequest),
// 			"status_code": fiber.StatusBadRequest,
// 			"message":     err.Error(),
// 			"result":      nil,
// 		})
// 	}

// 	var (
// 		ctx, cancel = context.WithTimeout(c.Request().Context(), time.Duration(30*time.Second))
// 		payload     dtos.CreateUserRequest
// 	)

// 	res, err := con.userUse.OnUserLogin(req)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"status":      http.StatusText(fiber.StatusNotFound),
// 			"status_code": fiber.StatusNotFound,
// 			"message":     err.Error(),
// 			"result":      nil,
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"status":      http.StatusText(fiber.StatusOK),
// 		"status_code": fiber.StatusOK,
// 		"message":     "",
// 		"result":      res,
// 	})
// }

func (con *userCon) FetchUsers(c *fiber.Ctx) error {

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	users, status, err := con.userUse.OnFetchUsers(ctx)

	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      http.StatusText(status),
			"status_code": status,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"status":      http.StatusText(status),
		"status_code": status,
		"message":     "",
		"result":      users,
	})
}

// TODO: continue the implementation of the FetchUserById function when finishing auth model
func (con *userCon) FetchUserById(c *fiber.Ctx) error {
	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
		reqP        = c.Get("userId")
	)
	defer cancel()

	if reqP == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "userId params is required",
			"result":      nil,
		})
	}

	userId, err := strconv.ParseInt(reqP, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	users, status, err := con.userUse.OnFetchUserById(ctx, userId)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      http.StatusText(status),
			"status_code": status,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"status":      http.StatusText(status),
		"status_code": status,
		"message":     "",
		"result":      users,
	})
}
