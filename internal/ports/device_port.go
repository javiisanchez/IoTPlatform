package ports

import (
	"IoTPlatform/internal/domain"

	fiber "github.com/gofiber/fiber/v2"
)

type IDeviceService interface {
	Create(device domain.Device) error
	Delete(ID string) error
}

type IDeviceRepository interface {
	Create(device domain.Device) error
	Delete(ID string) error
}

type IDeviceHandlers interface {
	Create(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type IServer interface {
	Initialize()
}
