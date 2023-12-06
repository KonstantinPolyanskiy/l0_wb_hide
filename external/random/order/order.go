package order

import (
	"fmt"
	"github.com/google/uuid"
	"l0_wb_hide/external/random/phone_number"
	"l0_wb_hide/external/random/rand_string"
	"l0_wb_hide/internal/models"
	"math/rand"
	"time"
)

// New возвращает Order, со случайно сгенерированными полями.
func New() models.Order {
	return models.Order{
		OrderUID:    uuid.New().String(),
		TrackNumber: "WBIL",
		Entry:       "WBIL",
		Delivery: models.Delivery{
			Name:    fmt.Sprintf("%s %s", rand_string.New(4), rand_string.New(6)),
			Phone:   phone_number.New(),
			Zip:     fmt.Sprintf("%s", rand.Intn(100000)),
			City:    rand_string.New(10),
			Address: rand_string.New(10),
			Region:  rand_string.New(7),
			Email:   rand_string.New(15) + "@mail.ru",
		},
		Payment: models.Payment{
			Transaction:  uuid.New().String() + "test",
			RequestId:    "",
			Currency:     "USD",
			Provider:     "wbpay",
			Amount:       rand.Intn(10000),
			PaymentDt:    rand.Int63n(100000000),
			Bank:         "alpha",
			DeliveryCost: rand.Intn(10000),
			GoodsTotal:   rand.Intn(1000),
			CustomFee:    0,
		},
		Items: []models.Item{{
			ChartId:     rand.Intn(10000000),
			TrackNumber: "WBILMTESTTRACK",
			Price:       rand.Intn(1000),
			Name:        rand_string.New(5),
			Rid:         uuid.New().String(),
			Sale:        rand.Intn(100),
			Size:        "0",
			TotalPrice:  rand.Intn(1000),
			NmId:        rand.Intn(10000000),
			Brand:       rand_string.New(9),
			Status:      202,
		}, {
			ChartId:     rand.Intn(10000000),
			TrackNumber: "WBILMTESTTRACK",
			Price:       rand.Intn(1000),
			Name:        rand_string.New(5),
			Rid:         uuid.New().String(),
			Sale:        rand.Intn(100),
			Size:        "0",
			TotalPrice:  rand.Intn(1000),
			NmId:        rand.Intn(10000000),
			Brand:       rand_string.New(9),
			Status:      202,
		}},
		Locale:            "en",
		InternalSignature: "",
		CustomerId:        "test",
		DeliveryService:   "meest",
		ShardKey:          "9",
		SmId:              rand.Intn(100),
		DateCreated:       time.Now(),
		OofChard:          "1",
	}
}
