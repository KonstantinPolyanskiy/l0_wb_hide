package main

import (
	"l0_wb_hide/external/http_server"
	"l0_wb_hide/internal/broker/stream"
	"l0_wb_hide/internal/cache"
	"l0_wb_hide/internal/handlers"
	"l0_wb_hide/internal/storage"
	"l0_wb_hide/internal/usecase"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cache := cache.New(
		cache.WithCapacity(10),
	)
	stream, err := stream.New(
		"test-cluster",
		"subscriber-1",
		"order_channel",
		"order-processing-service")
	if err != nil {
		log.Fatalf("Ошибка в запуске стрима - %s\n", err)
	}
	defer stream.Close()

	repository := storage.New(nil)
	service := usecase.New(repository, cache)
	handler := handlers.New(service, stream)

	server := http_server.New(handler.Init())

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.Run(); err != nil {
			log.Fatalf(err.Error())
		}
	}()
	<-done
}
