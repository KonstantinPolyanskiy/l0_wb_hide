package main

import (
	"l0_wb_hide/external/random/order"
	"l0_wb_hide/internal/broker/stream"
	"l0_wb_hide/internal/models"
	"log"
	"time"
)

const countMsg = 100

func main() {
	const op = "pub/main.go"

	str, err := stream.New(
		"test-cluster",
		"publisher-1",
		"order_channel",
		"order-processing-service")
	if err != nil {
		log.Fatalf("Ошибка с созданием stream %s в %s\n", err, op)
	}

	for i := 0; i <= 100; i++ {
		o := makeOrder()
		err = str.PublishOrder(o)
		if err != nil {
			log.Printf("Ошибка в отправке сообщения %s в %s\n", err, op)
			continue
		}
		log.Printf("Сообщение номер %d отправлено\n", i)

		time.Sleep(2 * time.Second)
	}

	defer str.Close()
}

func makeOrder() models.Order {
	return order.New()
}
