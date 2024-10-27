package mapper

import (
	"encoding/json"
	"fmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/valueobject"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/infrastructure/persistence/gorm/po"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/es"
	"github.com/stackus/errors"
)

type BasketMapperIntf interface {
	ToPersistent(basket *aggregate.Basket) (*po.Basket, error)
	ToDomain(basket *po.Basket) (*aggregate.Basket, error)
	ToDomainList(baskets []*po.Basket) ([]*aggregate.Basket, error)
}

type BasketMapper struct{}

var _ BasketMapperIntf = (*BasketMapper)(nil)

func NewBasketMapper() *BasketMapper {
	return &BasketMapper{}
}

func (b *BasketMapper) ToPersistent(basket *aggregate.Basket) (*po.Basket, error) {
	byt, err := json.Marshal(basket.Items)
	if err != nil {
		return nil, errors.ErrInternalServerError.Err(err)
	}

	return &po.Basket{
		ID:         basket.ID(),
		CustomerID: basket.CustomerID,
		PaymentID:  basket.PaymentID,
		Items:      byt,
		Status:     basket.Status.String(),
	}, nil
}

func (b *BasketMapper) ToDomain(basket *po.Basket) (*aggregate.Basket, error) {
	var items map[string]*entity.Item
	err := json.Unmarshal(basket.Items, &items)
	if err != nil {
		return nil, errors.ErrInternalServerError.Err(err)
	}

	status, err := b.statusToDomain(basket.Status)
	if err != nil {
		return nil, errors.ErrInternalServerError.Err(err)
	}

	return &aggregate.Basket{
		AggregateBase: es.NewAggregateBase(basket.ID, aggregate.BasketAggregate),
		CustomerID:    basket.CustomerID,
		PaymentID:     basket.PaymentID,
		Items:         items,
		Status:        status,
	}, nil
}

func (b *BasketMapper) ToDomainList(baskets []*po.Basket) ([]*aggregate.Basket, error) {
	rlt := make([]*aggregate.Basket, 0, len(baskets))

	for _, basket := range baskets {
		basketDO, err := b.ToDomain(basket)
		if err != nil {
			return nil, err
		}
		rlt = append(rlt, basketDO)
	}

	return rlt, nil
}

func (b *BasketMapper) statusToDomain(status string) (valueobject.BasketStatus, error) {
	switch status {
	case valueobject.BasketIsOpen.String():
		return valueobject.BasketIsOpen, nil
	case valueobject.BasketIsCancelled.String():
		return valueobject.BasketIsCancelled, nil
	case valueobject.BasketIsCheckedOut.String():
		return valueobject.BasketIsCheckedOut, nil
	default:
		return valueobject.BasketUnknown, fmt.Errorf("unknown basket status: %s", status)
	}
}
