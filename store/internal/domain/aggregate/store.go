package aggregate

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/event"
	"github.com/stackus/errors"
)

var (
	ErrStoreNameIsBlank               = errors.Wrap(errors.ErrBadRequest, "the store name cannot be blank")
	ErrStoreLocationIsBlank           = errors.Wrap(errors.ErrBadRequest, "the store location cannot be blank")
	ErrStoreIsAlreadyParticipating    = errors.Wrap(errors.ErrBadRequest, "the store is already participating")
	ErrStoreIsAlreadyNotParticipating = errors.Wrap(errors.ErrBadRequest, "the store is already not participating")
)

type Store struct {
	ddd.AggregateBase
	Name          string
	Location      string
	Participating bool
}

func CreateStore(id, name, location string) (*Store, error) {
	if name == "" {
		return nil, ErrStoreNameIsBlank
	}

	if location == "" {
		return nil, ErrStoreLocationIsBlank
	}

	store := &Store{
		AggregateBase: ddd.NewAggregateBase(id),
		Name:          name,
		Location:      location,
	}
	store.AddEvent(event.NewStoreCreatedEvent(name, location))

	return store, nil
}

func (store *Store) EnableParticipation() error {
	if store.Participating {
		return ErrStoreIsAlreadyParticipating
	}

	store.Participating = true
	store.AddEvent(event.NewStoreParticipationEnabledEvent())

	return nil
}

func (store *Store) DisableParticipation() error {
	if !store.Participating {
		return ErrStoreIsAlreadyNotParticipating
	}

	store.Participating = false
	store.AddEvent(event.NewStoreParticipationDisabledEvent())

	return nil
}
