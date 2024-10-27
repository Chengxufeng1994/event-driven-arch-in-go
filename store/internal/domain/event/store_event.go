package event

const (
	StoreCreatedEvent               = "stores.StoreCreated"
	StoreParticipationEnabledEvent  = "stores.StoreParticipationEnabled"
	StoreParticipationDisabledEvent = "stores.StoreParticipationDisabled"
	StoreReBrandedEvent             = "stores.StoreReBranded"
)

type StoreCreated struct {
	Name     string
	Location string
}

func NewStoreCreatedEvent(name, location string) *StoreCreated {
	return &StoreCreated{
		Name:     name,
		Location: location,
	}
}

func (StoreCreated) EventName() string { return StoreCreatedEvent }

type StoreParticipationDisabled struct {
	Participating bool
}

func NewStoreParticipationDisabledEvent() *StoreParticipationDisabled {
	return &StoreParticipationDisabled{Participating: false}
}

func (StoreParticipationDisabled) EventName() string { return StoreParticipationDisabledEvent }

type StoreParticipationEnabled struct {
	Participating bool
}

func NewStoreParticipationEnabledEvent() *StoreParticipationEnabled {
	return &StoreParticipationEnabled{Participating: true}
}

func (StoreParticipationEnabled) EventName() string { return StoreParticipationEnabledEvent }
