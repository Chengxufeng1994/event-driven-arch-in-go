package aggregate

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/valueobject"
)

type BasketV1 struct {
	CustomerID string
	PaymentID  string
	Items      map[string]*entity.Item
	Status     valueobject.BasketStatus
}

func (BasketV1) SnapshotName() string { return "baskets.BasketV1" }
