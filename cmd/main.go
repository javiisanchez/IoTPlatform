package main

import (
	"IoTPlatform/internal/adapters/api"
	"IoTPlatform/internal/adapters/repository"
	"IoTPlatform/internal/application"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Levantando IoTPlatform")

	repo := &repository.DeviceRepositoryImpl{}
	service := application.NewDeviceService(repo)
	handler := api.NewRestHandler(service)

	http.HandleFunc("/devices", handler.RegisterDevice)
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Printf("Servidor ejecutándose en el puerto 8080")

}
