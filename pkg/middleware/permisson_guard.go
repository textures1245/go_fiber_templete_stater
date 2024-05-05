package middleware

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/textures1245/go-template/pkg/utils"
)

func PermissionGuard(opt ...[]interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqP := c.Params("user_id")
		if reqP == "" {
			uuidBind := struct {
				UserId string `json:"user_id" form:"user_id" binding:"required"`
			}{}
			log.Info("user_id param not found, now trying to bind from request body")
			if err := c.BodyParser(&uuidBind); err != nil {
				return c.Status(http.StatusBadRequest).JSON(fiber.Map{
					"status":      http.StatusText(http.StatusBadRequest),
					"status_code": http.StatusBadRequest,
					"message":     "user_id is required",
					"raw_message": err.Error(),
					"result":      nil,
				})
			}
			reqP = uuidBind.UserId
		}
		uuidC, _ := c.Locals("user_id").(float64)
		userId, _ := strconv.Atoi(reqP)
		log.Debug(uuidC)

		if userId != int(uuidC) {
			if len(opt) > 0 {
				opt := opt[0]
				// if utils.Contains(opt, "PREVENT_DEFAULT_ACTION") {
				// 	return c.Next()
				// }
				log.Debug("opt: ", opt)
				if utils.Contains(opt, "PREVENT_DEFAULT_ACTION") {
					return c.Next()
				}
			}
			return c.Status(http.StatusForbidden).JSON(fiber.Map{
				"status":      http.StatusText(http.StatusForbidden),
				"status_code": http.StatusForbidden,
				"message":     "You don't have permission to access this resource",
				"raw_message": "",
				"result":      nil,
			})
		} else {
			if len(opt) > 0 {
				opt := opt[0]
				if utils.Contains(opt, "OWNER_ACTION_FORBIDDEN") {
					return c.Status(http.StatusForbidden).JSON(fiber.Map{
						"status":      http.StatusText(http.StatusForbidden),
						"status_code": http.StatusForbidden,
						"message":     "Owner action is forbidden",
						"raw_message": "",
						"result":      nil,
					})

				}
			}
			return c.Next()
		}
	}
}
