package gorm

import (
	"context"
	"fmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/sagastore/gorm/model"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/sec"
	"gorm.io/gorm"
)

type SagaStore struct {
	tableName string
	db        *gorm.DB
	registry  registry.Registry
}

var _ sec.SagaStore = (*SagaStore)(nil)

func NewSagaStore(tableName string, db *gorm.DB, registry registry.Registry) *SagaStore {
	return &SagaStore{
		tableName: tableName,
		db:        db,
		registry:  registry,
	}
}

// Load implements sec.SagaStore.
func (s *SagaStore) Load(ctx context.Context, sagaName string, sagaID string) (*sec.SagaContext[[]byte], error) {
	var saga model.Saga
	result := s.db.WithContext(ctx).
		Table(s.tableName).
		Select("data, step, done, compensating").
		Where(&model.Saga{ID: sagaID, Name: sagaName}).
		First(&saga)

	err := result.Error
	if err != nil {
		return nil, err
	}

	return &sec.SagaContext[[]byte]{
			ID:           sagaID,
			Data:         saga.Data,
			Step:         saga.Step,
			Done:         saga.Done,
			Compensating: saga.Compensating,
		},
		nil
}

// Save implements sec.SagaStore.
func (s *SagaStore) Save(ctx context.Context, sagaName string, sagaCtx *sec.SagaContext[[]byte]) error {
	const query = `INSERT INTO %s (name, id, data, step, done, compensating)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (name, id) DO
UPDATE SET data = EXCLUDED.data, step = EXCLUDED.step, done = EXCLUDED.done, compensating = EXCLUDED.compensating`
	stmt := fmt.Sprintf(query, s.tableName)

	err := s.db.WithContext(ctx).
		Exec(stmt, sagaName, sagaCtx.ID, sagaCtx.Data, sagaCtx.Step, sagaCtx.Done, sagaCtx.Compensating).Error
	if err != nil {
		return err
	}

	return nil
}
