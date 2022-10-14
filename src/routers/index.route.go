package routers

import "github.com/gofiber/fiber/v2"

func Index(route *fiber.App) {
	route.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"hello": "world!",
		})
	})
}
