package event

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
)

type MallDomainEvent interface {
	OnStoreCreated(ctx context.Context, event ddd.AggregateEvent) error
	OnStoreParticipationEnabled(ctx context.Context, event ddd.AggregateEvent) error
	OnStoreParticipationDisabled(ctx context.Context, event ddd.AggregateEvent) error
	OnStoreRebranded(ctx context.Context, event ddd.AggregateEvent) error
}
