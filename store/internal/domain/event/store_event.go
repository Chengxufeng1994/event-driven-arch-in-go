package event

const (
	StoreCreatedEvent               = "stores.StoreCreated"
	StoreParticipationEnabledEvent  = "stores.StoreParticipationEnabled"
	StoreParticipationDisabledEvent = "stores.StoreParticipationDisabled"
	StoreRebrandedEvent             = "stores.StoreRebranded"
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

func (StoreCreated) Key() string { return StoreCreatedEvent }

type StoreParticipationToggled struct {
	Participating bool
}

func NewStoreParticipationToggled(participating bool) *StoreParticipationToggled {
	return &StoreParticipationToggled{Participating: participating}
}

func (StoreParticipationToggled) Key() string { return StoreParticipationEnabledEvent }

type StoreRebranded struct {
	Name string
}

func NewStoreRebranded(name string) *StoreRebranded {
	return &StoreRebranded{Name: name}
}

func (StoreRebranded) Key() string { return StoreRebrandedEvent }
