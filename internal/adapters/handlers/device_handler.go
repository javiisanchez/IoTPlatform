package handlers

import (
	"IoTPlatform/internal/domain"
	"IoTPlatform/internal/ports"
	"encoding/json"
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
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

	// Device struct to be created from the request body
	var device domain.Device

	req := c.BodyRaw()
	fmt.Println(string(req))
	//test
	//var test = []byte(`{"ID":"ECA_0001","Name":"EstacionCalidadDelAire","Location":"BCN","Status":"True","CreatedAt":"2024-06-02T13:19:41Z","UpdatedAt":"2024-06-02T13:19:41Z"}`)
	//fmt.Println(string(test))

	err := json.Unmarshal(req, &device)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse request body")
	}
	fmt.Printf("%+v", device)

	//Extract the body and get the ID
	err1 := h.deviceService.Create(device)
	if err1 != nil {
		return err1
	}
	//repositories.NewDeviceRepository().CreateDevice(fiber.Ctx, device)
	return nil
}

func (h *DeviceHandlers) Delete(c *fiber.Ctx) error {
	var ID string
	//Extract the body and get the ID
	err := h.deviceService.Delete(ID)
	if err != nil {
		return err
	}
	return nil
}
