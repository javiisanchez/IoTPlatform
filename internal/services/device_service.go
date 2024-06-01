package services

import (
	"IoTPlatform/internal/ports"
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

func (s *DeviceService) Create(ID string) error {
	err := s.deviceRepository.Create(ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *DeviceService) Delete(ID string) error {

	err := s.deviceRepository.Delete(ID)
	if err != nil {
		return err
	}
	return nil
}
