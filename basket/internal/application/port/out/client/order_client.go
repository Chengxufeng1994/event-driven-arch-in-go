package client

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"
)

type OrderClient interface {
	Save(ctx context.Context, paymentID, customerID string, basketItems []*entity.Item) (string, error)
}
