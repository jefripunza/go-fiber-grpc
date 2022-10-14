package routers

import (
	"github.com/gofiber/fiber/v2"

	_ "go-fiber-grpc/docs"

	swagger "github.com/arsmn/fiber-swagger/v2"
)

/*
Refresh Swagger : swag init
*/

// SwaggerRoute func for describe group of API Docs routes.
func Swagger(route *fiber.App) {
	// Routes for GET method:
	route.Get("/swagger/*", swagger.HandlerDefault)
}
