package main

import (
	/*"IoTPlatform/internal/adapters/api"
	"IoTPlatform/internal/adapters/cli"
	"IoTPlatform/internal/application"
	"fmt"
	"log"
	"net/http"*/

	"IoTPlatform/internal/adapters/handlers"
	"IoTPlatform/internal/adapters/repositories"
	"IoTPlatform/internal/adapters/server"
	"IoTPlatform/internal/services"
	"fmt"
)

func main() {
	fmt.Println("Levantando IoTPlatform")

	mongoConn := "secret🤫"
	//repositories
	deviceRepository := repositories.NewDeviceRepository(mongoConn)
	//services
	deviceService := services.NewDeviceService(deviceRepository)
	//handlers
	deviceHandlers := handlers.NewDeviceHandlers(deviceService)
	//server
	s := server.NewServer(deviceHandlers)
	s.Initialize()
	//server.Initialize(userHandlers,)
	fmt.Printf("Servidor ejecutándose en el puerto 8080")
}
