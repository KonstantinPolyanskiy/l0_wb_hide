package stream

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	_ "github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"l0_wb_hide/internal/models"
	"log"
)

type Stream struct {
	nats        *nats.Conn
	stan        stan.Conn
	chanName    string
	durableName string
}

func New(clusterName, clientName, chanName, durableName string) (Stream, error) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return Stream{}, err
	}

	sc, err := stan.Connect(clusterName, clientName, stan.NatsConn(nc))
	if err != nil {
		return Stream{}, err
	}

	return Stream{
		nats:        nc,
		stan:        sc,
		chanName:    chanName,
		durableName: durableName,
	}, nil
}

func (s Stream) Close() {
	s.stan.Close()
	s.nats.Close()
}
func (s Stream) PublishOrder(order models.Order) error {
	data, err := json.Marshal(order)
	if err != nil {
		return err
	}

	err = s.stan.Publish(s.chanName, data)
	if err != nil {
		return err
	}

	return nil
}

func (s Stream) TakeOrder() (models.Order, error) {
	orderChan := make(chan models.Order, 1)

	sub, err := s.stan.Subscribe(s.chanName, func(msg *stan.Msg) {
		var order models.Order

		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			log.Println(err)
		} else {
			orderChan <- order
		}

	}, stan.DurableName(s.durableName))
	if err != nil {
		return models.Order{}, err
	}

	defer sub.Close()
	order := <-orderChan
	return order, nil
}
