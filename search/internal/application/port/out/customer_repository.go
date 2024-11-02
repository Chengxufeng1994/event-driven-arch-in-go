package out

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/domain"
)

type CustomerRepository interface {
	Find(ctx context.Context, customerID string) (*domain.Customer, error)
}
