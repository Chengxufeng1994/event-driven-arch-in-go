package entity

import "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/valueobject"

type Stops map[string]*Stop

func NewStops() Stops {
	return make(map[string]*Stop)
}

type Stop struct {
	Items         valueobject.Items
	StoreName     string
	StoreLocation string
}

func NewStop(storeName, storeLocation string) *Stop {
	return &Stop{
		Items:         valueobject.NewItems(),
		StoreName:     storeName,
		StoreLocation: storeLocation,
	}
}

func (s *Stop) AddItem(product valueobject.Product, quantity int) error {
	if _, exists := s.Items[product.ID]; !exists {
		item := valueobject.NewItem(product.Name, quantity)
		s.Items[product.ID] = &item
	}

	return nil
}
