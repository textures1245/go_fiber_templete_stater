package v1

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/textures1245/go-template/internal/user"
	"github.com/textures1245/go-template/internal/user/entities"
	"github.com/textures1245/go-template/pkg/utils"
)

type userCon struct {
	userUse user.UserUsecase
}

func NewUserController(userUse user.UserUsecase) *userCon {
	return &userCon{
		userUse: userUse,
	}
}

func (con userCon) UpdateUserById(c *fiber.Ctx) error {

	var req = new(entities.UserUpdateReq)
	if cE := c.BodyParser(req); cE != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid request body",
			"raw_message": cE.Error(),
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
		reqP        = c.Params("user_id")
	)
	defer cancel()

	if reqP == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "user_id params is required",
			"raw_message": "",
			"result":      nil,
		})
	}

	userId, err := strconv.ParseInt(reqP, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"raw_message": err.Error(),
			"message":     "error, invalid user_id params",
			"result":      nil,
		})
	}

	status, cE := con.userUse.OnUpdateUserById(ctx, userId, req)
	if cE != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      http.StatusText(status),
			"status_code": status,
			"message":     cE.CError.Error(),
			"raw_message": cE.RawError.Error(),
			"result":      nil,
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"status":      http.StatusText(status),
		"status_code": status,
		"message":     "user updated successfully",
		"result":      nil,
	})

}

func (con *userCon) FetchUsers(c *fiber.Ctx) error {

	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
	)

	defer cancel()

	users, status, cE := con.userUse.OnFetchUsers(ctx)

	if cE != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      http.StatusText(status),
			"status_code": status,
			"message":     cE.CError.Error(),
			"raw_message": cE.RawError.Error(),
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

func (con *userCon) FetchUserById(c *fiber.Ctx) error {
	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
		reqP        = c.Params("user_id")
	)
	defer cancel()

	if reqP == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "user_id params is required",
			"raw_message": "",
			"result":      nil,
		})
	}

	userId, err := strconv.ParseInt(reqP, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "",
			"raw_message": err.Error(),
			"result":      nil,
		})
	}

	users, status, cE := con.userUse.OnFetchUserById(ctx, userId)
	if cE != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      http.StatusText(status),
			"status_code": status,
			"message":     cE.CError.Error(),
			"raw_message": cE.RawError.Error(),
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

func (con *userCon) DeleteUserById(c *fiber.Ctx) error {
	var (
		ctx, cancel = context.WithTimeout(c.Context(), time.Duration(30*time.Second))
		reqP        = c.Params("user_id")
	)
	defer cancel()

	if reqP == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "userId params is required",
			"raw_message": "",
			"result":      nil,
		})
	}

	userId, err := strconv.ParseInt(reqP, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid userId params",
			"raw_message": err.Error(),
			"result":      nil,
		})
	}

	status, cE := con.userUse.UserDeleted(ctx, userId)
	if cE != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":      http.StatusText(status),
			"status_code": status,
			"message":     cE.CError.Error(),
			"raw_message": cE.RawError.Error(),
			"result":      nil,
		})
	}

	return c.Status(status).JSON(fiber.Map{
		"status":      http.StatusText(status),
		"status_code": status,
		"message":     "user deleted successfully",
		"result":      nil,
	})
}
