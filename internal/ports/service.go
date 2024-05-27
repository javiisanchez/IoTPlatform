package ports

import "IoTPlatform/internal/domain"

type DeviceService interface {
	RegisterDevice(device domain.Device) error
	GetDeviceReadings(deviceID string) ([]domain.Reading, error)
	//DeleteDeviceReadings(deviceID string) ([]domain.Reading, error)
	AddDeviceReadings(deviceID string) ([]domain.Reading, error)
}
