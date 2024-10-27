package ddd

const (
	AggregateNameKey    = "aggregate-name"
	AggregateIDKey      = "aggregate-id"
	AggregateVersionKey = "aggregate-version"
)

type Eventer interface {
	AddEvent(string, EventPayload, ...EventOption)
	Events() []AggregateEvent
	ClearEvents()
}

// AggregateEvent
type (
	AggregateEvent interface {
		Event
		AggregateName() string
		AggregateID() string
		AggregateVersion() int
	}

	aggregateEventBase struct {
		eventBase
	}
)

var _ AggregateEvent = (*aggregateEventBase)(nil)

func NewAggregateEventBase(name string, payload EventPayload, options ...EventOption) *aggregateEventBase {
	return &aggregateEventBase{
		eventBase: newEventBase(name, payload, options...),
	}
}

func (a *aggregateEventBase) AggregateID() string   { return a.metadata.Get(AggregateIDKey).(string) }
func (a *aggregateEventBase) AggregateName() string { return a.metadata.Get(AggregateNameKey).(string) }
func (a *aggregateEventBase) AggregateVersion() int { return a.metadata.Get(AggregateVersionKey).(int) }

// Aggregate
type AggregateNamer interface {
	AggregateName() string
}

type Aggregate interface {
	AggregateNamer
	Eventer
}

type AggregateBase struct {
	EntityBase
	events []AggregateEvent
}

var _ Aggregate = (*AggregateBase)(nil)

func NewAggregateBase(id, name string) AggregateBase {
	return AggregateBase{
		EntityBase: NewEntityBase(id, name),
		events:     make([]AggregateEvent, 0),
	}
}

func (agg *AggregateBase) AggregateName() string { return agg.name }
func (agg *AggregateBase) AddEvent(name string, payload EventPayload, options ...EventOption) {
	options = append(
		options,
		Metadata{
			AggregateNameKey: agg.name,
			AggregateIDKey:   agg.id,
		})

	agg.events = append(agg.events, NewAggregateEventBase(name, payload, options...))
}
func (agg *AggregateBase) Events() []AggregateEvent          { return agg.events }
func (agg *AggregateBase) ClearEvents()                      { agg.events = []AggregateEvent{} }
func (agg *AggregateBase) setEvents(events []AggregateEvent) { agg.events = events }
