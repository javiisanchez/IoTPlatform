package ports

import (
	"IoTPlatform/internal/domain"
	"context"

	fiber "github.com/gofiber/fiber/v2"
)

type IDeviceService interface {
	Create(device domain.Device) error
	Delete(ID string) error
}

type IDeviceRepository interface {
	Create(device domain.Device) error
	Delete(ID string) error
	CreateDevice(ctx context.Context, device *domain.Device) (*domain.Device, error)
}

type IRedingRepository interface {
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
