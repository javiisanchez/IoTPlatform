package ports

import "github.com/gofiber/fiber"

type IDeviceService interface {
	Create(ID string) error
	Delete(ID string) error
}

type IDeviceRepository interface {
	Create(ID string) error
	Delete(ID string) error
}

type IDeviceHandlers interface {
	Create(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type IServer interface {
	Initialize()
}
