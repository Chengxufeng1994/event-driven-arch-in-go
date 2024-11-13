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
	es.Aggregate
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
		Aggregate: es.NewAggregate(id, StoreAggregate),
	}
}

func (s *Store) InitStore(name, location string) (ddd.Event, error) {
	if name == "" {
		return nil, ErrStoreNameIsBlank
	}

	if location == "" {
		return nil, ErrStoreLocationIsBlank
	}

	s.AddEvent(event.StoreCreatedEvent, event.NewStoreCreatedEvent(
		name,
		location,
	))

	return ddd.NewEvent(event.StoreCreatedEvent, s), nil
}

// register serialize deserialize
func (Store) Key() string { return StoreAggregate }

func (s *Store) EnableParticipation() (ddd.Event, error) {
	if s.Participating {
		return nil, ErrStoreIsAlreadyParticipating
	}

	s.Participating = true

	s.AddEvent(event.StoreParticipationEnabledEvent, event.NewStoreParticipationToggled(
		true,
	))

	return ddd.NewEvent(event.StoreParticipationEnabledEvent, s), nil
}

func (s *Store) DisableParticipation() (ddd.Event, error) {
	if !s.Participating {
		return nil, ErrStoreIsAlreadyNotParticipating
	}

	s.Participating = false

	s.AddEvent(event.StoreParticipationDisabledEvent, event.NewStoreParticipationToggled(
		false,
	))

	return ddd.NewEvent(event.StoreParticipationDisabledEvent, s), nil
}

func (s *Store) Rebrand(name string) (ddd.Event, error) {
	s.AddEvent(event.StoreRebrandedEvent, event.NewStoreRebranded(
		name,
	))

	return ddd.NewEvent(event.StoreRebrandedEvent, s), nil
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
