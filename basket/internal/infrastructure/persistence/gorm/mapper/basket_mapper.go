package mapper

import (
	"encoding/json"
	"fmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/valueobject"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/infrastructure/persistence/gorm/po"
	"github.com/stackus/errors"
)

type BasketMapperIntf interface {
	ToPersistent(basket *aggregate.BasketAgg) (*po.Basket, error)
	ToDomain(basket *po.Basket) (*aggregate.BasketAgg, error)
	ToDomainList(baskets []*po.Basket) ([]*aggregate.BasketAgg, error)
}

type BasketMapper struct{}

var _ BasketMapperIntf = (*BasketMapper)(nil)

func NewBasketMapper() *BasketMapper {
	return &BasketMapper{}
}

func (b *BasketMapper) ToPersistent(basket *aggregate.BasketAgg) (*po.Basket, error) {
	byt, err := json.Marshal(basket.Items)
	if err != nil {
		return nil, errors.ErrInternalServerError.Err(err)
	}

	return &po.Basket{
		ID:         basket.ID,
		CustomerID: basket.CustomerID,
		PaymentID:  basket.PaymentID,
		Items:      byt,
		Status:     basket.Status.String(),
	}, nil
}

func (b *BasketMapper) ToDomain(basket *po.Basket) (*aggregate.BasketAgg, error) {
	var items []*entity.Item
	err := json.Unmarshal(basket.Items, &items)
	if err != nil {
		return nil, errors.ErrInternalServerError.Err(err)
	}

	status, err := b.statusToDomain(basket.Status)
	if err != nil {
		return nil, errors.ErrInternalServerError.Err(err)
	}

	return &aggregate.BasketAgg{
		ID:         basket.ID,
		CustomerID: basket.CustomerID,
		PaymentID:  basket.PaymentID,
		Items:      items,
		Status:     status,
	}, nil
}

func (b *BasketMapper) ToDomainList(baskets []*po.Basket) ([]*aggregate.BasketAgg, error) {
	rlt := make([]*aggregate.BasketAgg, 0, len(baskets))

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
	case valueobject.BasketOpen.String():
		return valueobject.BasketOpen, nil
	case valueobject.BasketCancelled.String():
		return valueobject.BasketCancelled, nil
	case valueobject.BasketCheckedOut.String():
		return valueobject.BasketCheckedOut, nil
	default:
		return valueobject.BasketUnknown, fmt.Errorf("unknown basket status: %s", status)
	}
}
