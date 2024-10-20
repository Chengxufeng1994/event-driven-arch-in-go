package entity

import "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/valueobject"

type Stops map[string]*Stop

func NewStops() Stops {
	return make(map[string]*Stop)
}

type Stop struct {
	StoreName     string
	StoreLocation string
	Items         valueobject.Items
}

func NewStop(storeName, storeLocation string) *Stop {
	return &Stop{
		StoreName:     storeName,
		StoreLocation: storeLocation,
		Items:         valueobject.NewItems(),
	}
}

func (s *Stop) AddItem(product valueobject.Product, quantity int) error {
	if _, exists := s.Items[product.ID]; !exists {
		item := valueobject.NewItem(product.Name, quantity)
		s.Items[product.ID] = &item
	}

	return nil
}
