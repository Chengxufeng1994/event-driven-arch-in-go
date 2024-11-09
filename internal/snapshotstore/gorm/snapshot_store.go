package gorm

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/es"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"gorm.io/gorm"
)

type SnapshotStore struct {
	es.AggregateStore
	db        *gorm.DB
	tableName string
	registry  registry.Registry
}

var _ es.AggregateStore = (*SnapshotStore)(nil)

func NewSnapshotStore(tableName string, db *gorm.DB, registry registry.Registry) es.AggregateStoreMiddleware {
	snapshots := SnapshotStore{
		tableName: tableName,
		db:        db,
		registry:  registry,
	}

	return func(store es.AggregateStore) es.AggregateStore {
		snapshots.AggregateStore = store
		return &snapshots
	}
}

func (s *SnapshotStore) Load(ctx context.Context, aggregate es.EventSourcedAggregate) error {
	const prepareStmt = `SELECT stream_version, snapshot_name, snapshot_data FROM %s WHERE stream_id = $1 AND stream_name = $2 LIMIT 1`
	stmt := fmt.Sprintf(prepareStmt, s.table())

	var entityVersion int
	var snapshotName string
	var snapshotData []byte

	err := s.db.WithContext(ctx).Raw(stmt, aggregate.ID(), aggregate.AggregateName()).Row().Scan(&entityVersion, &snapshotName, &snapshotData)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return s.AggregateStore.Load(ctx, aggregate)
		}
		return err
	}

	v, err := s.registry.Deserialize(snapshotName, snapshotData, registry.ValidateImplements((*es.Snapshot)(nil)))
	if err != nil {
		return err
	}

	if err := es.LoadSnapshot(aggregate, v.(es.Snapshot), entityVersion); err != nil {
		return err
	}

	return s.AggregateStore.Load(ctx, aggregate)
}

func (s *SnapshotStore) Save(ctx context.Context, aggregate es.EventSourcedAggregate) error {
	const prepareStmt = `INSERT INTO %s (stream_id, stream_name, stream_version, snapshot_name, snapshot_data)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (stream_id, stream_name) DO
UPDATE SET stream_version = EXCLUDED.stream_version, snapshot_name = EXCLUDED.snapshot_name, snapshot_data = EXCLUDED.snapshot_data
	`
	stmt := fmt.Sprintf(prepareStmt, s.table())

	if err := s.AggregateStore.Save(ctx, aggregate); err != nil {
		return err
	}

	if !s.shouldSnapshot(aggregate) {
		return nil
	}

	sser, ok := aggregate.(es.Snapshotter)
	if !ok {
		return fmt.Errorf("%T does not implelement es.Snapshotter", aggregate)
	}

	snapshot := sser.ToSnapshot()
	data, err := s.registry.Serialize(snapshot.SnapshotName(), snapshot)
	if err != nil {
		return err
	}

	err = s.db.WithContext(ctx).Exec(stmt, aggregate.ID(), aggregate.AggregateName(), aggregate.PendingVersion(), snapshot.SnapshotName(), data).Error
	if err != nil {
		return err
	}

	return nil
}

// TODO use injected & configurable strategies
func (SnapshotStore) shouldSnapshot(aggregate es.EventSourcedAggregate) bool {
	var maxChanges = 3 // low for demonstration; production envs should use higher values 50, 75, 100...
	var pendingVersion = aggregate.PendingVersion()
	var pendingChanges = len(aggregate.Events())

	return pendingVersion >= maxChanges && ((pendingChanges >= maxChanges) ||
		(pendingVersion%maxChanges < pendingChanges) ||
		(pendingVersion%maxChanges == 0))
}

func (s *SnapshotStore) table() string {
	return s.tableName
}
