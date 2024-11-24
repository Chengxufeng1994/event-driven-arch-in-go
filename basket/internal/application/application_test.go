package application

import (
	"context"
	"fmt"
	"testing"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/valueobject"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/es"
	"github.com/stretchr/testify/mock"
)

func TestApplication_AddItem(t *testing.T) {
	product := &entity.Product{
		ID:      "product-id",
		StoreID: "store-id",
		Name:    "product-name",
		Price:   10.00,
	}
	store := &entity.Store{
		ID:   "store-id",
		Name: "store-name",
	}

	type mocks struct {
		baskets   *repository.MockBasketRepository
		stores    *repository.MockStoreRepository
		products  *repository.MockProductRepository
		publisher *ddd.MockEventPublisher[ddd.Event]
	}
	type args struct {
		ctx context.Context
		add command.AddItem
	}
	tests := map[string]struct {
		args    args
		on      func(f mocks)
		wantErr bool
	}{
		"Success": {
			args: args{
				ctx: context.Background(),
				add: command.AddItem{
					ID:        "basket-id",
					ProductID: "product-id",
					Quantity:  1,
				},
			},
			on: func(f mocks) {
				f.baskets.On("Load", context.Background(), "basket-id").Return(&aggregate.Basket{
					Aggregate:  es.NewAggregate("basket-id", aggregate.BasketAggregate),
					CustomerID: "customer-id",
					PaymentID:  "payment-id",
					Items:      make(map[string]*entity.Item),
					Status:     valueobject.BasketIsOpen,
				}, nil)
				f.products.On("Find", context.Background(), "product-id").Return(product, nil)
				f.stores.On("Find", context.Background(), "store-id").Return(store, nil)
				f.baskets.On("Save", context.Background(), mock.AnythingOfType("*aggregate.Basket")).Return(nil)
			},
			wantErr: false,
		},
		"NoBasket": {
			args: args{
				ctx: context.Background(),
				add: command.AddItem{
					ID:        "basket-id",
					ProductID: "product-id",
					Quantity:  1,
				},
			},
			on: func(f mocks) {
				f.baskets.On("Load", context.Background(), "basket-id").Return(nil, fmt.Errorf("no basket"))
			},
			wantErr: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			m := mocks{
				baskets:   &repository.MockBasketRepository{},
				stores:    &repository.MockStoreRepository{},
				products:  &repository.MockProductRepository{},
				publisher: &ddd.MockEventPublisher[ddd.Event]{},
			}
			app := New(m.baskets, m.stores, m.products, m.publisher)
			if tt.on != nil {
				tt.on(m)
			}

			err := app.AddItem(tt.args.ctx, tt.args.add)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
