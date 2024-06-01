package server

import (
	"IoTPlatform/internal/ports"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	//We will add every new Handler here
	deviceHandlers ports.IDeviceHandlers
	//middlewares ports.IMiddlewares
	//paymentHandlers ports.IPaymentHandlers
}

func NewServer(uHandlers ports.IDeviceHandlers) *Server {
	return &Server{
		deviceHandlers: uHandlers,
		//paymentHandlers: pHandlers
	}
}

func (s *Server) Initialize() {
	app := fiber.New()
	v1 := app.Group("/v1")

	deviceRoutes := v1.Group("/user")
	deviceRoutes.Post("/login", func(c *fiber.Ctx) error {
		return c.SendString("I'm a POST request!")
	})
	//s.deviceHandlers.Create(c *fiber.Ctx)

	deviceRoutes.Post("/register", func(c *fiber.Ctx) error {
		return c.SendString("I'm a POST request!")
	})
	//s.deviceHandlers.Delete())

	err := app.Listen(":5000")
	if err != nil {
		log.Fatal(err)
	}
}
