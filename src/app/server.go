package app

import (
	"fmt"
	"go-fiber-grpc/src/configs"
	"log"

	"github.com/gofiber/fiber/v2"
)

// StartServer func for starting a simple server.
func WebServer(app *fiber.App) {

	// Run server.
	go func() {
		log.Fatal(app.Listen(fmt.Sprintf("0.0.0.0:%v", configs.GetPortWebServer())))
	}()

}
