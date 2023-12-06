package main

import (
	"l0_wb_hide/external/random/order"
	"l0_wb_hide/internal/broker/stream"
	"l0_wb_hide/internal/models"
	"log"
)

func main() {
	const op = "pub/main.go"

	stream, err := stream.New(
		"test-cluster",
		"publisher-1",
		"order_channel",
		"order-processing-service")
	if err != nil {
		log.Fatalf("Ошибка с созданием stream %s в %s\n", err, op)
	}
}

func makeOrder() models.Order {
	return order.New()
}
