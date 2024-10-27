package ddd

type Aggregate interface {
	Entity
	AddEvent(event DomainEvent)
	GetEvents() []DomainEvent
}

type AggregateBase struct {
	EntityBase
	events []DomainEvent
}

var _ Aggregate = (*AggregateBase)(nil)

func NewAggregateBase(id string) AggregateBase {
	return AggregateBase{
		EntityBase: NewEntityBase(id),
		events:     make([]DomainEvent, 0),
	}
}

func (agg *AggregateBase) GetID() string {
	return agg.ID
}

func (agg *AggregateBase) AddEvent(event DomainEvent) {
	agg.events = append(agg.events, event)
}

func (agg *AggregateBase) GetEvents() []DomainEvent {
	return agg.events
}
