package client

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification/internal/domain/valueobject"
)

type CustomerRepository interface {
	Find(ctx context.Context, customerID string) (valueobject.Customer, error)
}
