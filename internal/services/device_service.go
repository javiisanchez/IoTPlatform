package services

import (
	"IoTPlatform/internal/domain"
	"IoTPlatform/internal/ports"
	"fmt"
)

type DeviceService struct {
	deviceRepository ports.IDeviceRepository
	//deviceService    ports.IDeviceService
}

// This line is for get feedback in case we are not implementing the interface correctly
var _ ports.IDeviceService = (*DeviceService)(nil)

func NewDeviceService(repository ports.IDeviceRepository) *DeviceService {
	return &DeviceService{
		deviceRepository: repository,
	}
}

// Create a new device
func (s *DeviceService) Create(device domain.Device) error {
	err := s.deviceRepository.Create(device)
	if err != nil {
		return err
	}
	fmt.Println("Device created")
	return nil
}

func (s *DeviceService) Delete(ID string) error {

	err := s.deviceRepository.Delete(ID)
	if err != nil {
		return err
	}
	return nil
}
