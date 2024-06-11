package server

import (
	"IoTPlatform/internal/ports"
	"log"

	fiber "github.com/gofiber/fiber/v2"
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

	deviceRoutes.Post("/registerDevice", s.deviceHandlers.Create)

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
