package v1

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/textures1245/go-template/internal/user"
	"github.com/textures1245/go-template/internal/user/entities"
	"github.com/textures1245/go-template/pkg/apperror"
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

	var req = new(entities.UserUpdateReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "error, invalid request body",
			"raw_message": err.Error(),
			"result":      nil,
		})
	}

	status, err := con.userUse.OnUpdateUserById(ctx, userId, req)
	if err != nil {
		status, cE := apperror.CustomSqlExecuteHandler("User", err)
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
		reqP        = c.Params("user_id")
	)
	defer cancel()

	if reqP == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
			"message":     "user_id params is required",
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
		status, cE := apperror.CustomSqlExecuteHandler("User", err)
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

	status, err := con.userUse.UserDeleted(ctx, userId)
	if err != nil {
		status, cE := apperror.CustomSqlExecuteHandler("User", err)
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
