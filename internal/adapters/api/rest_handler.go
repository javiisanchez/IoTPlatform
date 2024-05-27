// internal/adapters/api/rest_handler.go
package api

import (
	"IoTPlatform/internal/application"
	"net/http"
)

type RestHandler struct {
	service *application.DeviceService
}

func NewRestHandler(service *application.DeviceService) *RestHandler {
	return &RestHandler{service: service}
}

func (h *RestHandler) RegisterDevice(w http.ResponseWriter, r *http.Request) {
	// Implementación del handler para registrar un dispositivo via REST
}

// Similarmente, implementamos handlers para gRPC, MQTT y Kafka
