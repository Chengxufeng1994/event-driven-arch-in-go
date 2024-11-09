package gorm

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/es"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"gorm.io/gorm"
)

type (
	EventStore struct {
		db        *gorm.DB
		tableName string
		registry  registry.Registry
	}

	aggregateEvent struct {
		id         string
		name       string
		payload    ddd.EventPayload
		occurredAt time.Time
		aggregate  es.EventSourcedAggregate
		version    int
	}
)

var _ es.AggregateStore = (*EventStore)(nil)

var _ ddd.AggregateEvent = (*aggregateEvent)(nil)

func NewEventStore(tableName string, db *gorm.DB, registry registry.Registry) *EventStore {
	return &EventStore{
		db:        db,
		tableName: tableName,
		registry:  registry,
	}
}

func (s EventStore) Load(ctx context.Context, aggregate es.EventSourcedAggregate) error {
	const prepareStmt = `SELECT stream_version, event_id, event_name, event_data, occurred_at FROM %s WHERE stream_id = $1 AND stream_name = $2 AND stream_version > $3 ORDER BY stream_version ASC`
	stmt := fmt.Sprintf(prepareStmt, s.table())

	aggregateID := aggregate.ID()
	aggregateName := aggregate.AggregateName()
	aggregateVersion := aggregate.Version()

	rows, err := s.db.WithContext(ctx).Raw(stmt,
		aggregateID,
		aggregateName,
		aggregateVersion,
	).Rows()
	if err != nil {
		return err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	for rows.Next() {
		var eventID, eventName string
		var payloadData []byte
		var aggregateVersion int
		var occurredAt time.Time
		err := rows.Scan(&aggregateVersion, &eventID, &eventName, &payloadData, &occurredAt)
		if err != nil {
			return err
		}

		payload, err := s.registry.Deserialize(eventName, payloadData)
		if err != nil {
			return err
		}

		event := aggregateEvent{
			id:         eventID,
			name:       eventName,
			payload:    payload,
			aggregate:  aggregate,
			version:    aggregateVersion,
			occurredAt: occurredAt,
		}

		if err := es.LoadEvent(aggregate, event); err != nil {
			return err
		}
	}

	return nil
}

func (s EventStore) Save(ctx context.Context, aggregate es.EventSourcedAggregate) (err error) {
	const prepareStmt = `INSERT INTO %s (stream_id, stream_name, stream_version, event_id, event_name, event_data, occurred_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	stmt := fmt.Sprintf(prepareStmt, s.table())

	aggregateID := aggregate.ID()
	aggregateName := aggregate.AggregateName()

	for _, event := range aggregate.Events() {
		payload, err := s.registry.Serialize(event.EventName(), event.Payload())
		if err != nil {
			return err
		}

		err = s.db.WithContext(ctx).Exec(stmt,
			aggregateID,
			aggregateName,
			event.AggregateVersion(),
			event.ID(),
			event.EventName(),
			payload,
			event.OccurredAt(),
		).Error

		if err != nil {
			return err
		}
	}

	return nil
}

func (s EventStore) table() string {
	return s.tableName
}

func (e aggregateEvent) ID() string                { return e.id }
func (e aggregateEvent) EventName() string         { return e.name }
func (e aggregateEvent) Payload() ddd.EventPayload { return e.payload }
func (e aggregateEvent) Metadata() ddd.Metadata    { return ddd.Metadata{} }
func (e aggregateEvent) OccurredAt() time.Time     { return e.occurredAt }
func (e aggregateEvent) AggregateName() string     { return e.aggregate.AggregateName() }
func (e aggregateEvent) AggregateID() string       { return e.aggregate.ID() }
func (e aggregateEvent) AggregateVersion() int     { return e.version }
