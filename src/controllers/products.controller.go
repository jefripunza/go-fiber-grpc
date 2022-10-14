package controllers

import (
	"fmt"
	"go-fiber-grpc/src/dto/request"
	"go-fiber-grpc/src/handlers"
	"go-fiber-grpc/src/messages"
	"go-fiber-grpc/src/models/repositories"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ProductCreateOne(c *fiber.Ctx) error {
	// Check, if received JSON data is valid.
	dto := &request.CreateProduct{}
	if err := c.BodyParser(dto); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var result error
	handlers.Error{
		Try: func() {
			repositories.InsertProduct(dto)
			result = c.JSON(fiber.Map{
				"message": messages.SaveDataSuccess,
			})
		},
		Catch: func(e handlers.Exception) {
			fmt.Printf("Caught error : %v\n", e)
			result = c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": e,
			})
		},
		Finally: func() {},
	}.Do()
	return result
}

func ProductReadAllData(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": repositories.ReadProducts(),
	})
}

func ProductById(c *fiber.Ctx) error {
	// ambil parameter dari endpoint
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		err_msg := err.Error()
		if strings.HasPrefix(err_msg, "strconv.Atoi") {
			err_msg = "id harus angka!"
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err_msg,
		})
	}

	var result error
	handlers.Error{
		Try: func() {
			result = c.JSON(repositories.ReadProductById(id))
		},
		Catch: func(e handlers.Exception) {
			fmt.Printf("Caught error : %v\n", e)
			result = c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": fmt.Sprintf("%v", e),
			})
		},
		Finally: func() {},
	}.Do()
	return result
}

func ProductByCode(c *fiber.Ctx) error {
	// ambil parameter dari endpoint
	code := c.Params("code")

	var result error
	handlers.Error{
		Try: func() {
			result = c.JSON(repositories.ReadProductByCode(code))
		},
		Catch: func(e handlers.Exception) {
			fmt.Printf("Caught error : %v\n", e)
			result = c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": fmt.Sprintf("%v", e),
			})
		},
		Finally: func() {},
	}.Do()
	return result
}

func ProductUpdateById(c *fiber.Ctx) error {
	// Check, if received JSON data is valid.
	dto := &request.UpdateProduct{}
	if err := c.BodyParser(dto); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var result error
	handlers.Error{
		Try: func() {
			repositories.UpdateProduct(dto)
			result = c.JSON(fiber.Map{
				"message": messages.UpdateDataSuccess,
			})
		},
		Catch: func(e handlers.Exception) {
			fmt.Printf("Caught error : %v\n", e)
			result = c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": e,
			})
		},
		Finally: func() {},
	}.Do()
	return result
}

func ProductDeleteById(c *fiber.Ctx) error {
	// ambil parameter dari endpoint
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		err_msg := err.Error()
		if strings.HasPrefix(err_msg, "strconv.Atoi") {
			err_msg = "id harus angka!"
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err_msg,
		})
	}

	var result error
	handlers.Error{
		Try: func() {
			repositories.DeleteProductById(id)
			result = c.JSON(fiber.Map{
				"message": messages.DeleteDataSuccess,
			})
		},
		Catch: func(e handlers.Exception) {
			fmt.Printf("Caught error : %v\n", e)
			result = c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": e,
			})
		},
		Finally: func() {},
	}.Do()
	return result
}
