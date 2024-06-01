package handlers

import (
	"IoTPlatform/internal/ports"

	"github.com/gofiber/fiber"
)

type DeviceHandlers struct {
	deviceService ports.IDeviceService
}

var _ ports.IDeviceHandlers = (*DeviceHandlers)(nil)

func NewDeviceHandlers(deviceService ports.IDeviceService) *DeviceHandlers {
	return &DeviceHandlers{
		deviceService: deviceService,
	}
}

func (h *DeviceHandlers) Create(c *fiber.Ctx) error {
	var ID string
	//Extract the body and get the email and password
	err := h.deviceService.Create(ID)
	if err != nil {
		return err
	}
	return nil
}

func (h *DeviceHandlers) Delete(c *fiber.Ctx) error {
	var ID string
	//Extract the body and get the email and password
	err := h.deviceService.Delete(ID)
	if err != nil {
		return err
	}
	return nil
}
