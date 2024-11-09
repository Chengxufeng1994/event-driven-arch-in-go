package gorm

import (
	"context"
	"fmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
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
	sagaCtx := &sec.SagaContext[[]byte]{
		ID: sagaID,
	}

	row := s.db.WithContext(ctx).
		Table(s.tableName).
		Select("data, step, done, compensating").
		Where("name = ? AND id = ?", sagaName, sagaID).
		Row()

	err := row.Scan(&sagaCtx.Data, &sagaCtx.Step, &sagaCtx.Done, &sagaCtx.Compensating)

	return sagaCtx, err
}

// Save implements sec.SagaStore.
func (s *SagaStore) Save(ctx context.Context, sagaName string, sagaCtx *sec.SagaContext[[]byte]) error {
	const query = `INSERT INTO %s (name, id, data, step, done, compensating) 
VALUES ($1, $2, $3, $4, $5, $6) 
ON CONFLICT (name, id) DO
UPDATE SET data = EXCLUDED.data, step = EXCLUDED.step, done = EXCLUDED.done, compensating = EXCLUDED.compensating`

	stmt := fmt.Sprintf(query, s.tableName)
	result := s.db.WithContext(ctx).
		Exec(stmt, sagaName, sagaCtx.ID, sagaCtx.Data, sagaCtx.Step, sagaCtx.Done, sagaCtx.Compensating)

	return result.Error
}
