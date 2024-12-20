package mapper

import (
	"encoding/json"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/valueobject"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/infastructure/persistence/gorm/po"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/stackus/errors"
)

type ShoppingListMapperIntf interface {
	ToPersistent(*aggregate.ShoppingList) (*po.ShoppingList, error)
	ToDomain(*po.ShoppingList) (*aggregate.ShoppingList, error)
}

type ShoppingListMapper struct{}

var _ ShoppingListMapperIntf = (*ShoppingListMapper)(nil)

func NewShoppingListMapper() ShoppingListMapperIntf {
	return &ShoppingListMapper{}
}

func (m *ShoppingListMapper) ToPersistent(shoppingList *aggregate.ShoppingList) (*po.ShoppingList, error) {
	byt, err := json.Marshal(shoppingList.Stops)
	if err != nil {
		return nil, errors.Wrap(err, "json marshal")
	}

	return &po.ShoppingList{
		ID:            shoppingList.ID(),
		OrderID:       shoppingList.OrderID,
		AssignedBotID: shoppingList.AssignedBotID,
		Stops:         byt,
		Status:        shoppingList.Status.String(),
	}, nil
}

func (m *ShoppingListMapper) ToDomain(shoppingList *po.ShoppingList) (*aggregate.ShoppingList, error) {
	stops := entity.NewStops()
	err := json.Unmarshal(shoppingList.Stops, &stops)
	if err != nil {
		return nil, errors.Wrap(err, "json unmarshal")
	}

	return &aggregate.ShoppingList{
		Aggregate:     ddd.NewAggregate(shoppingList.ID, aggregate.ShoppingListAggregate),
		OrderID:       shoppingList.OrderID,
		AssignedBotID: shoppingList.AssignedBotID,
		Stops:         stops,
		Status:        valueobject.NewShoppingListStatus(shoppingList.Status),
	}, nil
}

func (m *ShoppingListMapper) statusToDomain(status string) valueobject.ShoppingListStatus {
	switch status {
	case valueobject.ShoppingListIsPending.String():
		return valueobject.ShoppingListIsPending
	case valueobject.ShoppingListIsAvailable.String():
		return valueobject.ShoppingListIsAvailable
	case valueobject.ShoppingListIsAssigned.String():
		return valueobject.ShoppingListIsAssigned
	case valueobject.ShoppingListIsActive.String():
		return valueobject.ShoppingListIsActive
	case valueobject.ShoppingListIsCompleted.String():
		return valueobject.ShoppingListIsCompleted
	case valueobject.ShoppingListIsCanceled.String():
		return valueobject.ShoppingListIsCanceled
	default:
		return valueobject.ShoppingListUnknown
	}
}
