package mapper

import (
	"encoding/json"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/valueobject"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/infastructure/persistence/gorm/po"
	"github.com/stackus/errors"
)

type ShoppingListMapperIntf interface {
	ToPersistence(*aggregate.ShoppingListAgg) (*po.ShoppingList, error)
	ToDomain(*po.ShoppingList) (*aggregate.ShoppingListAgg, error)
}

type ShoppingListMapper struct{}

var _ ShoppingListMapperIntf = (*ShoppingListMapper)(nil)

func NewShoppingListMapper() ShoppingListMapperIntf {
	return &ShoppingListMapper{}
}

func (m *ShoppingListMapper) ToPersistence(shoppingList *aggregate.ShoppingListAgg) (*po.ShoppingList, error) {
	byt, err := json.Marshal(shoppingList.Stops)
	if err != nil {
		return nil, errors.Wrap(err, "json marshal")
	}

	return &po.ShoppingList{
		ID:            shoppingList.ID,
		OrderID:       shoppingList.OrderID,
		AssignedBotID: shoppingList.AssignedBotID,
		Stops:         byt,
		Status:        shoppingList.Status.String(),
	}, nil
}

func (m *ShoppingListMapper) ToDomain(shoppingList *po.ShoppingList) (*aggregate.ShoppingListAgg, error) {
	stops := entity.NewStops()
	err := json.Unmarshal(shoppingList.Stops, &stops)
	if err != nil {
		return nil, errors.Wrap(err, "json unmarshal")
	}

	return &aggregate.ShoppingListAgg{
		ID:            shoppingList.ID,
		OrderID:       shoppingList.OrderID,
		AssignedBotID: shoppingList.AssignedBotID,
		Stops:         stops,
		Status:        m.statusToDomain(shoppingList.Status),
	}, nil
}

func (m *ShoppingListMapper) statusToDomain(status string) valueobject.ShoppingListStatus {
	switch status {
	case valueobject.ShoppingListAvailable.String():
		return valueobject.ShoppingListAvailable
	case valueobject.ShoppingListAssigned.String():
		return valueobject.ShoppingListAssigned
	case valueobject.ShoppingListActive.String():
		return valueobject.ShoppingListActive
	case valueobject.ShoppingListCompleted.String():
		return valueobject.ShoppingListCompleted
	case valueobject.ShoppingListCancelled.String():
		return valueobject.ShoppingListCancelled
	default:
		return valueobject.ShoppingListUnknown
	}
}
