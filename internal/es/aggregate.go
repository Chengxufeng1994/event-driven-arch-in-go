package es

import "github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"

type (
	Versioner interface {
		Version() int
		PendingVersion() int
	}

	// event sourcing Aggregate base on ddd Aggregate
	Aggregate interface {
		ddd.Aggregate
		EventCommitter
		Versioner
		VersionSetter
	}

	AggregateBase struct {
		ddd.AggregateBase
		version int
	}
)

var _ Aggregate = (*AggregateBase)(nil)

func NewAggregateBase(id, name string) AggregateBase {
	return AggregateBase{
		AggregateBase: ddd.NewAggregateBase(id, name),
		version:       0,
	}
}

func (a *AggregateBase) AddEvent(name string, payload ddd.EventPayload, options ...ddd.EventOption) {
	options = append(
		options,
		ddd.Metadata{
			ddd.AggregateVersionKey: a.PendingVersion() + 1,
		})

	a.AggregateBase.AddEvent(name, payload, options...)
}

func (a *AggregateBase) CommitEvents() {
	a.version += len(a.Events())
	a.ClearEvents()
}

func (a *AggregateBase) Version() int           { return a.version }
func (a *AggregateBase) PendingVersion() int    { return a.version + len(a.Events()) }
func (a *AggregateBase) setVersion(version int) { a.version = version }
