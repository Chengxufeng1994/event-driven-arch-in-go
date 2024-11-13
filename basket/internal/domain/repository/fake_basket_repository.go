package repository

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/aggregate"
)

type FakeBasketRepository struct {
	baskets map[string]*aggregate.Basket
}

func NewFakeBasketRepository() *FakeBasketRepository {
	return &FakeBasketRepository{baskets: map[string]*aggregate.Basket{}}
}

var _ BasketRepository = (*FakeBasketRepository)(nil)

func (r *FakeBasketRepository) Load(ctx context.Context, basketID string) (*aggregate.Basket, error) {
	if basket, exists := r.baskets[basketID]; exists {
		return basket, nil
	}

	return aggregate.NewBasket(basketID), nil
}

func (r *FakeBasketRepository) Save(ctx context.Context, basket *aggregate.Basket) error {
	r.baskets[basket.ID()] = basket

	return nil
}

func (r *FakeBasketRepository) Reset(baskets ...*aggregate.Basket) {
	r.baskets = make(map[string]*aggregate.Basket)

	for _, basket := range baskets {
		r.baskets[basket.ID()] = basket
	}
}
