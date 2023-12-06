package stream

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	_ "github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"l0_wb_hide/internal/models"
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
