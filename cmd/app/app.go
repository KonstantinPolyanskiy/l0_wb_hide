package main

import (
	"l0_wb_hide/external/http_server"
	"l0_wb_hide/internal/cache"
	"l0_wb_hide/internal/handlers"
	"l0_wb_hide/internal/storage"
	"l0_wb_hide/internal/usecase"
	"log"
)

func main() {
	cache := cache.New(
		cache.WithCapacity(10),
	)
	repository := storage.New(nil)
	service := usecase.New(repository, cache)
	handler := handlers.New(service)

	server := http_server.New(handler.Init())

	go func() {
		log.Fatal(server.Run())
	}()
}
