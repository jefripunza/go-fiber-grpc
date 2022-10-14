package routers

import (
	"github.com/gofiber/fiber/v2"

	"go-fiber-grpc/src/controllers"
)

func Basic(r fiber.Router) {
	route_v1 := r.Group("/v1")
	route_v1.Get("/add/:a/:b", controllers.BasicAdd)
	route_v1.Get("/multiply/:a/:b", controllers.BasicMultiply)
}
