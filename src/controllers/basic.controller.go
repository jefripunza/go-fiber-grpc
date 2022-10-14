package controllers

import (
	"go-fiber-grpc/src/communications"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func BasicAdd(c *fiber.Ctx) error {
	// Check, if received JSON data is valid.
	a, err := strconv.ParseUint(c.Params("a"), 10, 64)
	if err != nil {
		err_msg := err.Error()
		if strings.HasPrefix(err_msg, "strconv.ParseUint") {
			err_msg = "a harus angka!"
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err_msg,
		})
	}
	b, err := strconv.ParseUint(c.Params("b"), 10, 64)
	if err != nil {
		err_msg := err.Error()
		if strings.HasPrefix(err_msg, "strconv.ParseUint") {
			err_msg = "b harus angka!"
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err_msg,
		})
	}

	result, err_msg := communications.BasicAdd(a, b)
	if err_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err_msg.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"result": result,
	})
}

func BasicMultiply(c *fiber.Ctx) error {
	// Check, if received JSON data is valid.
	a, err := strconv.ParseUint(c.Params("a"), 10, 64)
	if err != nil {
		err_msg := err.Error()
		if strings.HasPrefix(err_msg, "strconv.ParseUint") {
			err_msg = "a harus angka!"
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err_msg,
		})
	}
	b, err := strconv.ParseUint(c.Params("b"), 10, 64)
	if err != nil {
		err_msg := err.Error()
		if strings.HasPrefix(err_msg, "strconv.ParseUint") {
			err_msg = "b harus angka!"
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err_msg,
		})
	}

	result, err_msg := communications.BasicMultiply(a, b)
	if err_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err_msg.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"result": result,
	})
}
