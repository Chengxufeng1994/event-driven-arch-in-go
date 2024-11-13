package aggregate

import (
	"testing"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/valueobject"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/es"
	"github.com/stretchr/testify/assert"
)

func TestBasket_AddItem(t *testing.T) {
	store := &entity.Store{
		ID:   "store-id",
		Name: "store-name",
	}
	product := &entity.Product{
		ID:      "product-id",
		StoreID: "store-id",
		Name:    "product-name",
		Price:   10.00,
	}

	type fields struct {
		CustomerID string
		PaymentID  string
		Items      map[string]*entity.Item
		Status     valueobject.BasketStatus
	}
	type args struct {
		store    *entity.Store
		product  *entity.Product
		quantity int
	}
	tests := map[string]struct {
		fields  fields
		args    args
		on      func(a *es.MockAggregate)
		wantErr bool
	}{
		"OpenBasket": {
			fields: fields{
				Items:  make(map[string]*entity.Item),
				Status: valueobject.BasketIsOpen,
			},
			args: args{
				store:    store,
				product:  product,
				quantity: 1,
			},
			on: func(a *es.MockAggregate) {
				a.On("AddEvent", event.BasketItemAddedEvent, &event.BasketItemAdded{
					Item: &entity.Item{
						StoreID:      store.ID,
						ProductID:    product.ID,
						StoreName:    store.Name,
						ProductName:  product.Name,
						ProductPrice: product.Price,
						Quantity:     1,
					},
				})
			},
			wantErr: false,
		},
		"CheckedOutBasket": {
			fields: fields{
				Items:  make(map[string]*entity.Item),
				Status: valueobject.BasketIsCheckedOut,
			},
			args: args{
				store:    store,
				product:  product,
				quantity: 1,
			},
			wantErr: true,
		},
		"CanceledOutBasket": {
			fields: fields{
				Items:  make(map[string]*entity.Item),
				Status: valueobject.BasketIsCanceled,
			},
			args: args{
				store:    store,
				product:  product,
				quantity: 1,
			},
			wantErr: true,
		},
		"ZeroQuantity": {
			fields: fields{
				Items:  make(map[string]*entity.Item),
				Status: valueobject.BasketIsCheckedOut,
			},
			args: args{
				store:    store,
				product:  product,
				quantity: 0,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			aggregate := es.NewMockAggregate(t)
			b := &Basket{
				Aggregate:  aggregate,
				CustomerID: tt.fields.CustomerID,
				PaymentID:  tt.fields.PaymentID,
				Items:      tt.fields.Items,
				Status:     tt.fields.Status,
			}
			if tt.on != nil {
				tt.on(aggregate)
			}

			if err := b.AddItem(tt.args.store, tt.args.product, tt.args.quantity); (err != nil) != tt.wantErr {
				t.Errorf("AddItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBasket_ApplyEvent(t *testing.T) {
	store := entity.Store{
		ID:   "store-id",
		Name: "store-name",
	}
	product := entity.Product{
		ID:      "product-id",
		StoreID: "store-id",
		Name:    "product-name",
		Price:   10.00,
	}
	product2 := entity.Product{
		ID:      "product-id-2",
		StoreID: "store-id",
		Name:    "product-name-2",
		Price:   10.00,
	}
	type fields struct {
		CustomerID string
		PaymentID  string
		Items      map[string]*entity.Item
		Status     valueobject.BasketStatus
	}
	type args struct {
		event ddd.Event
	}

	tests := map[string]struct {
		fields  fields
		args    args
		want    fields
		wantErr bool
	}{
		"BasketItemAddEvent": {
			fields: fields{
				CustomerID: "customer-id",
				PaymentID:  "payment-id",
				Items:      make(map[string]*entity.Item),
				Status:     valueobject.BasketIsOpen,
			},
			args: args{
				event: ddd.NewEvent(event.BasketItemAddedEvent, &event.BasketItemAdded{
					Item: &entity.Item{
						StoreID:      store.ID,
						ProductID:    product.ID,
						StoreName:    store.Name,
						ProductName:  product.Name,
						ProductPrice: product.Price,
						Quantity:     1,
					},
				}),
			},
			want: fields{
				CustomerID: "customer-id",
				PaymentID:  "payment-id",
				Items: map[string]*entity.Item{
					product.ID: {
						StoreID:      store.ID,
						ProductID:    product.ID,
						StoreName:    store.Name,
						ProductName:  product.Name,
						ProductPrice: product.Price,
						Quantity:     1,
					},
				},
				Status: valueobject.BasketIsOpen,
			},
			wantErr: false,
		},
		"BasketItemAddedEvent.Quantity": {
			fields: fields{
				CustomerID: "customer-id",
				PaymentID:  "payment-id",
				Items: map[string]*entity.Item{
					product.ID: {
						StoreID:      store.ID,
						ProductID:    product.ID,
						StoreName:    store.Name,
						ProductName:  product.Name,
						ProductPrice: product.Price,
						Quantity:     1,
					},
				},
				Status: valueobject.BasketIsOpen,
			},
			args: args{
				event: ddd.NewEvent(event.BasketItemAddedEvent, &event.BasketItemAdded{
					Item: &entity.Item{
						StoreID:      store.ID,
						ProductID:    product.ID,
						StoreName:    store.Name,
						ProductName:  product.Name,
						ProductPrice: product.Price,
						Quantity:     1,
					},
				}),
			},
			want: fields{
				CustomerID: "customer-id",
				PaymentID:  "payment-id",
				Items: map[string]*entity.Item{
					product.ID: {
						StoreID:      store.ID,
						ProductID:    product.ID,
						StoreName:    store.Name,
						ProductName:  product.Name,
						ProductPrice: product.Price,
						Quantity:     2,
					},
				},
				Status: valueobject.BasketIsOpen,
			},
			wantErr: false,
		},
		"BasketItemAddedEvent.Second": {
			fields: fields{
				CustomerID: "customer-id",
				PaymentID:  "payment-id",
				Items: map[string]*entity.Item{
					product.ID: {
						StoreID:      store.ID,
						ProductID:    product.ID,
						StoreName:    store.Name,
						ProductName:  product.Name,
						ProductPrice: product.Price,
						Quantity:     1,
					},
				},
				Status: valueobject.BasketIsOpen,
			},
			args: args{
				event: ddd.NewEvent(event.BasketItemAddedEvent, &event.BasketItemAdded{
					Item: &entity.Item{
						StoreID:      store.ID,
						ProductID:    product2.ID,
						StoreName:    store.Name,
						ProductName:  product2.Name,
						ProductPrice: product2.Price,
						Quantity:     1,
					},
				}),
			},
			want: fields{
				CustomerID: "customer-id",
				PaymentID:  "payment-id",
				Items: map[string]*entity.Item{
					product.ID: {
						StoreID:      store.ID,
						ProductID:    product.ID,
						StoreName:    store.Name,
						ProductName:  product.Name,
						ProductPrice: product.Price,
						Quantity:     1,
					},
					product2.ID: {
						StoreID:      store.ID,
						ProductID:    product2.ID,
						StoreName:    store.Name,
						ProductName:  product2.Name,
						ProductPrice: product2.Price,
						Quantity:     1,
					},
				},
				Status: valueobject.BasketIsOpen,
			},
			wantErr: false,
		},
		"BasketCanceledEvent": {
			fields: fields{
				CustomerID: "customer-id",
				PaymentID:  "payment-id",
				Items:      make(map[string]*entity.Item),
				Status:     valueobject.BasketIsOpen,
			},
			args: args{event: ddd.NewEvent(event.BasketCanceledEvent, &event.BasketCanceled{})},
			want: fields{
				CustomerID: "customer-id",
				PaymentID:  "payment-id",
				Items:      map[string]*entity.Item{},
				Status:     valueobject.BasketIsCanceled,
			},
			wantErr: false,
		},
		"BasketCanceledEvent.Cleared": {
			fields: fields{
				CustomerID: "customer-id",
				PaymentID:  "payment-id",
				Items: map[string]*entity.Item{
					product.ID: {
						StoreID:      store.ID,
						ProductID:    product.ID,
						StoreName:    store.Name,
						ProductName:  product.Name,
						ProductPrice: product.Price,
						Quantity:     1,
					},
				},
				Status: valueobject.BasketIsOpen,
			},
			args: args{event: ddd.NewEvent(event.BasketCanceledEvent, &event.BasketCanceled{})},
			want: fields{
				CustomerID: "customer-id",
				PaymentID:  "payment-id",
				Items:      map[string]*entity.Item{},
				Status:     valueobject.BasketIsCanceled,
			},
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			b := &Basket{
				Aggregate:  es.NewMockAggregate(t),
				CustomerID: tt.fields.CustomerID,
				PaymentID:  tt.fields.PaymentID,
				Items:      tt.fields.Items,
				Status:     tt.fields.Status,
			}

			if err := b.ApplyEvent(tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("ApplyEvent() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.Equal(t, b.CustomerID, tt.want.CustomerID)
			assert.Equal(t, b.PaymentID, tt.want.PaymentID)
			assert.Equal(t, b.Items, tt.want.Items)
			assert.Equal(t, b.Status, tt.want.Status)
		})
	}
}

func TestBasket_ApplySnapshot(t *testing.T) {
	store := &entity.Store{
		ID:   "store-id",
		Name: "store-name",
	}
	product := &entity.Product{
		ID:      "product-id",
		StoreID: "store-id",
		Name:    "product-name",
		Price:   10.00,
	}
	item := &entity.Item{
		StoreID:     store.ID,
		ProductID:   product.ID,
		StoreName:   store.Name,
		ProductName: product.Name,
		Quantity:    1,
	}

	type fields struct {
		CustomerID string
		PaymentID  string
		Items      map[string]*entity.Item
		Status     valueobject.BasketStatus
	}

	type args struct {
		snapshot es.Snapshot
	}

	tests := map[string]struct {
		fields  fields
		args    args
		want    fields
		wantErr bool
	}{
		"V1": {
			fields: fields{},
			args: args{
				snapshot: &BasketV1{
					CustomerID: "customer-id",
					PaymentID:  "payment-id",
					Items:      map[string]*entity.Item{product.ID: item},
					Status:     valueobject.BasketIsOpen,
				},
			},
			want: fields{
				CustomerID: "customer-id",
				PaymentID:  "payment-id",
				Items:      map[string]*entity.Item{product.ID: item},
				Status:     valueobject.BasketIsOpen,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			b := &Basket{
				Aggregate:  es.NewMockAggregate(t),
				CustomerID: tt.fields.CustomerID,
				PaymentID:  tt.fields.PaymentID,
				Items:      tt.fields.Items,
				Status:     tt.fields.Status,
			}
			if err := b.ApplySnapshot(tt.args.snapshot); (err != nil) != tt.wantErr {
				t.Errorf("ApplySnapshot() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.Equal(t, b.CustomerID, tt.want.CustomerID)
			assert.Equal(t, b.PaymentID, tt.want.PaymentID)
			assert.Equal(t, b.Items, tt.want.Items)
			assert.Equal(t, b.Status, tt.want.Status)
		})
	}

}

func TestBasket_Cancel(t *testing.T) {
	type fields struct {
		CustomerID string
		PaymentID  string
		Items      map[string]*entity.Item
		Status     valueobject.BasketStatus
	}
	tests := map[string]struct {
		fields  fields
		on      func(a *es.MockAggregate)
		want    ddd.Event
		wantErr bool
	}{
		"OpenBasket": {
			fields: fields{
				CustomerID: "customer-id",
				PaymentID:  "payment-id",
				Items:      make(map[string]*entity.Item),
				Status:     valueobject.BasketIsOpen,
			},
			on: func(a *es.MockAggregate) {
				a.On("AddEvent", event.BasketCanceledEvent, &event.BasketCanceled{})
			},
			want: ddd.NewEvent(event.BasketCanceledEvent, &Basket{
				CustomerID: "customer-id",
				PaymentID:  "payment-id",
				Items:      make(map[string]*entity.Item),
				Status:     valueobject.BasketIsCanceled,
			}),
		},
		"CheckedOutBasket": {
			fields: fields{
				CustomerID: "customer-id",
				PaymentID:  "payment-id",
				Items:      make(map[string]*entity.Item),
				Status:     valueobject.BasketIsCheckedOut,
			},
			wantErr: true,
		},
		"CanceledBasket": {
			fields: fields{
				CustomerID: "customer-id",
				PaymentID:  "payment-id",
				Items:      make(map[string]*entity.Item),
				Status:     valueobject.BasketIsCanceled,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			aggregate := es.NewMockAggregate(t)
			b := &Basket{
				Aggregate:  aggregate,
				CustomerID: tt.fields.CustomerID,
				PaymentID:  tt.fields.PaymentID,
				Items:      tt.fields.Items,
				Status:     tt.fields.Status,
			}
			if tt.on != nil {
				tt.on(aggregate)
			}

			got, err := b.Cancel()
			if (err != nil) != tt.wantErr {
				t.Errorf("Cancel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				assert.Equal(t, tt.want.EventName(), got.EventName())
				assert.IsType(t, tt.want.Payload(), got.Payload())
				assert.Equal(t, tt.want.Metadata(), got.Metadata())
			} else {
				assert.Nil(t, got)
			}
		})
	}
}

func TestBasket_Checkout(t *testing.T) {
	store := &entity.Store{
		ID:   "store-id",
		Name: "store-name",
	}
	product := &entity.Product{
		ID:      "product-id",
		StoreID: "store-id",
		Name:    "product-name",
		Price:   10.00,
	}
	item := entity.Item{
		StoreID:      store.ID,
		ProductID:    product.ID,
		StoreName:    store.Name,
		ProductName:  product.Name,
		ProductPrice: product.Price,
		Quantity:     1,
	}

	type fields struct {
		CustomerID string
		PaymentID  string
		Items      map[string]*entity.Item
		Status     valueobject.BasketStatus
	}
	type args struct {
		paymentID string
	}
	tests := map[string]struct {
		fields  fields
		args    args
		on      func(a *es.MockAggregate)
		want    ddd.Event
		wantErr bool
	}{
		"OpenBasket": {
			fields: fields{
				CustomerID: "customer-id",
				PaymentID:  "payment-id",
				Items: map[string]*entity.Item{
					product.ID: &item,
				},
				Status: valueobject.BasketIsOpen,
			},
			args: args{paymentID: "payment-id"},
			on: func(a *es.MockAggregate) {
				a.On("AddEvent", event.BasketCheckedOutEvent, &event.BasketCheckedOut{
					CustomerID: "customer-id",
					PaymentID:  "payment-id",
					Items: map[string]*entity.Item{
						product.ID: &item,
					},
				})
			},
			want: ddd.NewEvent(event.BasketCheckedOutEvent, &Basket{
				CustomerID: "customer-id",
				PaymentID:  "payment-id",
				Items: map[string]*entity.Item{
					product.ID: &item,
				},
				Status: valueobject.BasketIsCheckedOut,
			}),
		},
		"OpenBasket.NoItems": {
			fields: fields{
				CustomerID: "customer-id",
				PaymentID:  "payment-id",
				Items:      make(map[string]*entity.Item),
				Status:     valueobject.BasketIsOpen,
			},
			args:    args{paymentID: "payment-id"},
			wantErr: true,
		},
		"CheckedOutBasket": {
			fields: fields{
				CustomerID: "customer-id",
				PaymentID:  "payment-id",
				Items:      make(map[string]*entity.Item),
				Status:     valueobject.BasketIsCheckedOut,
			},
			args:    args{paymentID: "payment-id"},
			wantErr: true,
		},
		"CanceledBasket": {
			fields: fields{
				CustomerID: "customer-id",
				PaymentID:  "payment-id",
				Items:      make(map[string]*entity.Item),
				Status:     valueobject.BasketIsCanceled,
			},
			args:    args{paymentID: "payment-id"},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// Arrange
			aggregate := es.NewMockAggregate(t)
			b := &Basket{
				Aggregate:  aggregate,
				CustomerID: tt.fields.CustomerID,
				PaymentID:  tt.fields.PaymentID,
				Items:      tt.fields.Items,
				Status:     tt.fields.Status,
			}
			if tt.on != nil {
				tt.on(aggregate)
			}

			// Act
			got, err := b.Checkout(tt.args.paymentID)

			// Assert
			if (err != nil) != tt.wantErr {
				t.Errorf("Checkout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				assert.Equal(t, tt.want.EventName(), got.EventName())
				assert.IsType(t, tt.want.Payload(), got.Payload())
				assert.Equal(t, tt.want.Metadata(), got.Metadata())
			} else {
				assert.Nil(t, got)
			}
		})
	}
}

func TestBasket_RemoveItem(t *testing.T) {
	store := &entity.Store{
		ID:   "store-id",
		Name: "store-name",
	}
	product := &entity.Product{
		ID:      "product-id",
		StoreID: "store-id",
		Name:    "product-name",
		Price:   10.00,
	}
	item := &entity.Item{
		StoreID:      store.ID,
		ProductID:    product.ID,
		StoreName:    store.Name,
		ProductName:  product.Name,
		ProductPrice: product.Price,
		Quantity:     10,
	}

	type fields struct {
		CustomerID string
		PaymentID  string
		Items      map[string]*entity.Item
		Status     valueobject.BasketStatus
	}
	type args struct {
		product  *entity.Product
		quantity int
	}
	tests := map[string]struct {
		fields  fields
		args    args
		on      func(a *es.MockAggregate)
		wantErr bool
	}{
		"OpenBasket": {
			fields: fields{
				CustomerID: "customer-id",
				PaymentID:  "payment-id",
				Items: map[string]*entity.Item{
					product.ID: item,
				},
				Status: valueobject.BasketIsOpen,
			},
			args: args{
				product:  product,
				quantity: 1,
			},
			on: func(a *es.MockAggregate) {
				a.On("AddEvent", event.BasketItemRemovedEvent, &event.BasketItemRemoved{
					ProductID: product.ID,
					Quantity:  1,
				})
			},
		},
		"OpenBasket.NoItems": {
			fields: fields{
				CustomerID: "customer-id",
				PaymentID:  "payment-id",
				Items:      make(map[string]*entity.Item),
				Status:     valueobject.BasketIsOpen,
			},
			args: args{
				product:  product,
				quantity: 1,
			},
		},
		"CheckedOutBasket": {
			fields: fields{
				CustomerID: "customer-id",
				PaymentID:  "payment-id",
				Items:      make(map[string]*entity.Item),
				Status:     valueobject.BasketIsCheckedOut,
			},
			args: args{
				product:  product,
				quantity: 1,
			},
			wantErr: true,
		},
		"CanceledBasket": {
			fields: fields{
				CustomerID: "customer-id",
				PaymentID:  "payment-id",
				Items:      make(map[string]*entity.Item),
				Status:     valueobject.BasketIsCanceled,
			},
			args: args{
				product:  product,
				quantity: 1,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			aggregate := es.NewMockAggregate(t)
			b := &Basket{
				Aggregate:  aggregate,
				CustomerID: tt.fields.CustomerID,
				PaymentID:  tt.fields.PaymentID,
				Items:      tt.fields.Items,
				Status:     tt.fields.Status,
			}
			if tt.on != nil {
				tt.on(aggregate)
			}

			if err := b.RemoveItem(tt.args.product, tt.args.quantity); (err != nil) != tt.wantErr {
				t.Errorf("RemoveItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBasket_Start(t *testing.T) {
	type fields struct {
		CustomerID string
		PaymentID  string
		Items      map[string]*entity.Item
		Status     valueobject.BasketStatus
	}
	type args struct {
		customerID string
	}
	tests := map[string]struct {
		fields  fields
		args    args
		on      func(a *es.MockAggregate)
		want    ddd.Event
		wantErr bool
	}{
		"New": {
			fields: fields{},
			args:   args{customerID: "customer-id"},
			on: func(a *es.MockAggregate) {
				a.On("AddEvent", event.BasketStartedEvent, &event.BasketStarted{
					CustomerID: "customer-id",
				})
			},
			want: ddd.NewEvent(event.BasketStartedEvent, &Basket{
				CustomerID: "customer-id",
				PaymentID:  "",
				Items:      make(map[string]*entity.Item),
				Status:     valueobject.BasketIsOpen,
			}),
			wantErr: false,
		},
		"OpenBasket": {
			fields: fields{
				Status: valueobject.BasketIsOpen,
			},
			args:    args{customerID: "customer-id"},
			wantErr: true,
		},
		"CheckedOutBasket": {
			fields: fields{
				Status: valueobject.BasketIsCheckedOut,
			},
			args:    args{customerID: "customer-id"},
			wantErr: true,
		},
		"CanceledBasket": {
			fields: fields{
				Status: valueobject.BasketIsCanceled,
			},
			args:    args{customerID: "customer-id"},
			wantErr: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			aggregate := es.NewMockAggregate(t)
			b := &Basket{
				Aggregate:  aggregate,
				CustomerID: tt.fields.CustomerID,
				PaymentID:  tt.fields.PaymentID,
				Items:      tt.fields.Items,
				Status:     tt.fields.Status,
			}
			if tt.on != nil {
				tt.on(aggregate)
			}

			got, err := b.Start(tt.args.customerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Start() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.want != nil {
				assert.Equal(t, tt.want.EventName(), got.EventName())
				assert.IsType(t, tt.want.Payload(), got.Payload())
			} else {
				assert.Nil(t, got)
			}
		})
	}
}

func TestNewBasket(t *testing.T) {
	type args struct {
		id string
	}
	tests := map[string]struct {
		args args
		want *Basket
	}{
		"Basket": {
			args: args{id: "basket-id"},
			want: &Basket{
				Aggregate: es.NewAggregate("basket-id", BasketAggregate),
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := NewBasket(tt.args.id)

			assert.Equal(t, tt.want.ID(), got.ID())
			assert.Equal(t, tt.want.AggregateName(), got.AggregateName())
		})
	}
}
