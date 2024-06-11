package main

import (
	"IoTPlatform/internal/adapters/config"
	"IoTPlatform/internal/adapters/handlers"
	"IoTPlatform/internal/adapters/repositories"
	"IoTPlatform/internal/adapters/server"
	"IoTPlatform/internal/adapters/storage/postgres"
	"IoTPlatform/internal/services"
	"context"
	"fmt"
	"log/slog"
	"os"
)

func main() {
	fmt.Println("Levantando IoTPlatform")

	// Load environment variables
	config, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	// Set logger
	//logger.Set(config.App)

	slog.Info("Starting the application", "app", config.App.Name, "env", config.App.Env)

	// Init database
	ctx := context.Background()
	db, err := postgres.New(ctx, config.DB)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	slog.Info("Successfully connected to the database", "db", config.DB.Connection)

	// Migrate database
	err = db.Migrate()
	if err != nil {
		slog.Error("Error migrating database", "error", err)
		os.Exit(1)
	}

	slog.Info("Successfully migrated the database")

	//repositories
	deviceRepository := repositories.NewDeviceRepository(db)
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
