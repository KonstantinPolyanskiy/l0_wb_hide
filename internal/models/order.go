package models

import (
	"fmt"
	"time"
)

var ErrNoItems = fmt.Errorf("no items")

type Order struct {
	OrderUID          string `json:"order_uid"`
	TrackNumber       string `json:"track_number"`
	Entry             string `json:"entry"`
	Delivery          Delivery
	Payment           Payment
	Items             []Item
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	ShardKey          string    `json:"shardkey"`
	SmId              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofChard          string    `json:"oof_chard"`
}
type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     int    `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}
type Payment struct {
	Transaction  string `json:"transaction"`
	RequestId    int    `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int64  `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type Item struct {
	ChartId     int    `json:"chart_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	TotalPrice  int    `json:"total_price"`
	NmId        int    `json:"nmId"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

func (o Order) Validate() error {
	if len(o.Items) <= 0 {
		return ErrNoItems
	}

	return nil
}
