package aggregate

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/es"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/event"
	"github.com/stackus/errors"
)

const StoreAggregate = "stores.Store"

var (
	ErrStoreNameIsBlank               = errors.Wrap(errors.ErrBadRequest, "the store name cannot be blank")
	ErrStoreLocationIsBlank           = errors.Wrap(errors.ErrBadRequest, "the store location cannot be blank")
	ErrStoreIsAlreadyParticipating    = errors.Wrap(errors.ErrBadRequest, "the store is already participating")
	ErrStoreIsAlreadyNotParticipating = errors.Wrap(errors.ErrBadRequest, "the store is already not participating")
)

type Store struct {
	es.AggregateBase
	Name          string
	Location      string
	Participating bool
}

var _ interface {
	es.EventApplier
	es.Snapshotter
} = (*Store)(nil)

func NewStore(id string) *Store {
	return &Store{
		AggregateBase: es.NewAggregateBase(id, StoreAggregate),
	}
}

func CreateStore(id, name, location string) (*Store, error) {
	if name == "" {
		return nil, ErrStoreNameIsBlank
	}

	if location == "" {
		return nil, ErrStoreLocationIsBlank
	}

	store := NewStore(id)

	store.AddEvent(event.StoreCreatedEvent, event.NewStoreCreatedEvent(name, location))

	return store, nil
}

// register serialize deserialize
func (Store) Key() string { return StoreAggregate }

func (store *Store) EnableParticipation() error {
	if store.Participating {
		return ErrStoreIsAlreadyParticipating
	}

	store.Participating = true

	store.AddEvent(event.StoreParticipationEnabledEvent, event.NewStoreParticipationToggled(
		true,
	))

	return nil
}

func (store *Store) DisableParticipation() error {
	if !store.Participating {
		return ErrStoreIsAlreadyNotParticipating
	}

	store.Participating = false

	store.AddEvent(event.StoreParticipationDisabledEvent, event.NewStoreParticipationToggled(
		false,
	))

	return nil
}

func (store *Store) Rebrand(name string) error {
	store.AddEvent(event.StoreRebrandedEvent, event.NewStoreRebranded(name))

	return nil
}

func (store *Store) ApplyEvent(e ddd.Event) error {
	switch payload := e.Payload().(type) {
	case *event.StoreCreated:
		store.Name = payload.Name
		store.Location = payload.Location

	case *event.StoreParticipationToggled:
		store.Participating = payload.Participating

	case *event.StoreRebranded:
		store.Name = payload.Name

	default:
		return errors.ErrInternal.Msgf("%T received the event %s with unexpected payload %T", store, e.EventName(), payload)
	}

	return nil
}

func (store *Store) ApplySnapshot(snapshot es.Snapshot) error {
	switch ss := snapshot.(type) {
	case *StoreV1:
		store.Name = ss.Name
		store.Location = ss.Location
		store.Participating = ss.Participating
	default:
		return errors.ErrInternal.Msgf("%T received the unexpected snapshot %T", store, snapshot)
	}

	return nil
}

func (store *Store) ToSnapshot() es.Snapshot {
	return StoreV1{
		Name:          store.Name,
		Location:      store.Location,
		Participating: store.Participating,
	}
}
