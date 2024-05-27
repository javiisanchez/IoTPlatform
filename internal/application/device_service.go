package application

import (
	"IoTPlatform/internal/domain"
	"IoTPlatform/internal/ports"
)

// Class
type DeviceService struct {
	repo ports.DeviceRepository
}

func NewDeviceService(repo ports.DeviceRepository) *DeviceService {
	return &DeviceService{repo: repo}
}

// Pasa el device para guardarlo en BBDD
func (s *DeviceService) RegisterDevice(device domain.Device) error {
	// Validaciones y lógica de negocio
	return s.repo.SaveDevice(device)
}

func (s *DeviceService) GetDeviceReadings(deviceID string) ([]domain.Reading, error) {
	// Obtener y procesar lecturas
	return nil, nil
}
