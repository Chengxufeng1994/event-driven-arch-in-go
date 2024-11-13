package ddd

const (
	AggregateNameKey    = "aggregate-name"
	AggregateIDKey      = "aggregate-id"
	AggregateVersionKey = "aggregate-version"
)

// Aggregate
type (
	AggregateNamer interface {
		AggregateName() string
	}

	Eventer interface {
		AddEvent(string, EventPayload, ...EventOption)
		Events() []AggregateEvent
		ClearEvents()
	}

	Aggregate interface {
		IDer
		AggregateNamer
		Eventer
		IDSetter
		NameSetter
	}

	aggregate struct {
		Entity
		events []AggregateEvent
	}
)

var _ Aggregate = (*aggregate)(nil)

func NewAggregate(id, name string) *aggregate {
	return &aggregate{
		Entity: NewEntity(id, name),
		events: make([]AggregateEvent, 0),
	}
}

func (agg aggregate) AggregateName() string    { return agg.EntityName() }
func (agg aggregate) Events() []AggregateEvent { return agg.events }
func (agg *aggregate) ClearEvents()            { agg.events = []AggregateEvent{} }
func (agg *aggregate) AddEvent(name string, payload EventPayload, options ...EventOption) {
	options = append(
		options,
		Metadata{
			AggregateNameKey: agg.EntityName(),
			AggregateIDKey:   agg.ID(),
		})
	agg.events = append(
		agg.events,
		NewAggregateEvent(name, payload, options...))
}

func (agg *aggregate) setEvents(events []AggregateEvent) { agg.events = events }

// AggregateEvent
type (
	AggregateEvent interface {
		Event
		AggregateName() string
		AggregateID() string
		AggregateVersion() int
	}

	aggregateEvent struct {
		event
	}
)

var _ AggregateEvent = (*aggregateEvent)(nil)

func NewAggregateEvent(name string, payload EventPayload, options ...EventOption) aggregateEvent {
	return aggregateEvent{
		event: newEvent(name, payload, options...),
	}
}

func (a aggregateEvent) AggregateID() string   { return a.metadata.Get(AggregateIDKey).(string) }
func (a aggregateEvent) AggregateName() string { return a.metadata.Get(AggregateNameKey).(string) }
func (a aggregateEvent) AggregateVersion() int { return a.metadata.Get(AggregateVersionKey).(int) }
