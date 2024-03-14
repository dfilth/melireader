package main

import (
	"fmt"
	"log/slog"
	"melireader/internal/adapter/config"
	"melireader/internal/adapter/handler/http"
	"melireader/internal/adapter/logger"
	"melireader/internal/adapter/storage/documentDb"
	"melireader/internal/adapter/storage/documentDb/repository"
	"melireader/internal/core/service"
	"os"
)

// @title					Go MELI FILE READER API
// @version					1.0
// @description				This is a simple RESTFUL Point of Sale (POS) Service API written in Go using
// Gin web framework and Mongo database with Hexagonal Architecture
// @contact.name			Ing. Daniel Torres
// @contact.url				https://github.com/
// @BasePath				/v1
// @schemes					http https
func main() {
	// Load environment variables
	cfg, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	logger.Set(cfg.App)

	slog.Info("Starting the application", "app", cfg.App.Name, "env", cfg.App.Env)

	db := documentDb.MeliDB(cfg.DB)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}

	slog.Info("Successfully connected to the database", "db", cfg.DB.Connection)

	// Dependency injection
	itemRepo := repository.NewItemRepository(db)
	itemService := service.NewItemService(itemRepo, nil)
	itemHandler := http.NewFileItemHandler(itemService)

	// Init router
	router, err := http.NewRouter(
		cfg.HTTP,
		*itemHandler,
	)
	if err != nil {
		slog.Error("Error initializing router", "error", err)
		os.Exit(1)
	}

	// Start server
	listenAddr := fmt.Sprintf("%s:%s", cfg.HTTP.URL, cfg.HTTP.Port)
	slog.Info("Starting the HTTP server", "listen_address", listenAddr)
	err = router.Serve(listenAddr)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}
