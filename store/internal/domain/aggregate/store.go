package aggregate

import "github.com/stackus/errors"

var (
	ErrStoreNameIsBlank               = errors.Wrap(errors.ErrBadRequest, "the store name cannot be blank")
	ErrStoreLocationIsBlank           = errors.Wrap(errors.ErrBadRequest, "the store location cannot be blank")
	ErrStoreIsAlreadyParticipating    = errors.Wrap(errors.ErrBadRequest, "the store is already participating")
	ErrStoreIsAlreadyNotParticipating = errors.Wrap(errors.ErrBadRequest, "the store is already not participating")
)

type StoreAgg struct {
	ID            string
	Name          string
	Location      string
	Participating bool
}

func CreateStore(id, name, location string) (*StoreAgg, error) {
	if name == "" {
		return nil, ErrStoreNameIsBlank
	}

	if location == "" {
		return nil, ErrStoreLocationIsBlank
	}

	store := &StoreAgg{
		ID:       id,
		Name:     name,
		Location: location,
	}

	return store, nil
}

func (store *StoreAgg) EnableParticipation() error {
	if store.Participating {
		return ErrStoreIsAlreadyParticipating
	}

	store.Participating = true

	return nil
}

func (store *StoreAgg) DisableParticipation() error {
	if !store.Participating {
		return ErrStoreIsAlreadyNotParticipating
	}

	store.Participating = false
	return nil
}
