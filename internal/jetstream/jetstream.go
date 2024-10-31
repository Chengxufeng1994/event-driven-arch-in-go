package jetstream

import (
	"fmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/config"
	"github.com/nats-io/nats.go"
)

func NewJetStream(config *config.Infrastructure, nc *nats.Conn) (nats.JetStreamContext, error) {
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}
	_, err = js.AddStream(&nats.StreamConfig{
		Name:     config.Nats.Stream,
		Subjects: []string{fmt.Sprintf("%s.>", config.Nats.Stream)},
	})

	return js, err
}
