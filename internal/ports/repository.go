package ports

import "IoTPlatform/internal/domain"

type DeviceRepository interface {
	SaveDevice(device domain.Device) error
	GetDevice(id string) (*domain.Device, error)
	//DeleteDevice(id string) (*domain.Device, error)
	AddDevice(device domain.Device) error
}
