package gorm

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/es"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/stackus/errors"
	"gorm.io/gorm"
)

type GormEventStore struct {
	db        *gorm.DB
	tableName string
	registry  registry.Registry
}

var _ es.AggregateStore = (*GormEventStore)(nil)

func NewGormEventStore(tableName string, db *gorm.DB, registry registry.Registry) *GormEventStore {
	return &GormEventStore{
		db:        db,
		tableName: tableName,
		registry:  registry,
	}
}

// Load implements es.AggregateStore.
func (s GormEventStore) Load(ctx context.Context, aggregate es.EventSourcedAggregate) error {
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

// Save implements es.AggregateStore.
func (s GormEventStore) Save(ctx context.Context, aggregate es.EventSourcedAggregate) (err error) {
	tx := s.db.WithContext(ctx).Begin(&sql.TxOptions{})
	defer func() {
		p := recover()
		switch {
		case p != nil:
			_ = tx.Rollback()
			panic(p)
		case tx.Error != nil:
			rErr := tx.Rollback().Error
			if rErr != nil {
				err = errors.Wrap(tx.Error, rErr.Error())
			}
		default:
			err = tx.Commit().Error
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	const prepareStmt = `INSERT INTO %s (stream_id, stream_name, stream_version, event_id, event_name, event_data, occurred_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	stmt := fmt.Sprintf(prepareStmt, s.table())

	aggregateID := aggregate.ID()
	aggregateName := aggregate.AggregateName()

	for _, event := range aggregate.Events() {
		payload, err := s.registry.Serialize(event.EventName(), event.Payload())
		if err != nil {
			return err
		}

		err = tx.Exec(stmt,
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

func (s GormEventStore) table() string {
	return s.tableName
}
