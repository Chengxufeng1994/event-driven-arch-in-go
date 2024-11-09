package gorm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/outboxstore/gorm/model"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/tm"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type OutboxStore struct {
	tableName string
	db        *gorm.DB
}

type outboxMessage struct {
	id      string
	name    string
	subject string
	data    []byte
}

var _ tm.OutboxStore = (*OutboxStore)(nil)
var _ am.RawMessage = (*outboxMessage)(nil)

func NewOutboxStore(tableName string, db *gorm.DB) *OutboxStore {
	return &OutboxStore{
		tableName: tableName,
		db:        db,
	}
}

// Save implements tm.OutboxStore.
func (s *OutboxStore) Save(ctx context.Context, msg am.RawMessage) error {
	const query = "INSERT INTO %s (id, name, subject, data) VALUES ($1, $2, $3, $4)"
	stmt := fmt.Sprintf(query, s.tableName)

	result := s.db.WithContext(ctx).
		Exec(stmt, msg.ID(), msg.MessageName(), msg.Subject(), msg.Data())

	if result.Error != nil {
		var pqErr *pq.Error
		if errors.As(result.Error, &pqErr) {
			if pqErr.Code == pgerrcode.UniqueViolation {
				return tm.ErrDuplicateMessage(msg.ID())
			}
		}
		return result.Error
	}

	return nil
}

// FindUnpublished implements tm.OutboxStore.
func (s *OutboxStore) FindUnpublished(ctx context.Context, limit int) ([]am.RawMessage, error) {
	var msgs []model.Outbox

	result := s.db.WithContext(ctx).
		Table(s.tableName).
		Where("published_at IS NULL").
		Limit(limit).
		Find(&msgs)

	if result.Error != nil {
		return nil, result.Error
	}

	outboxMsgs := make([]am.RawMessage, 0, len(msgs))
	for _, msg := range msgs {
		outboxMsgs = append(outboxMsgs, &outboxMessage{
			id:      msg.ID,
			name:    msg.Name,
			subject: msg.Subject,
			data:    msg.Data,
		})
	}

	return outboxMsgs, nil
}

// MarkPublished implements tm.OutboxStore.
func (s *OutboxStore) MarkPublished(ctx context.Context, ids ...string) error {
	result := s.db.WithContext(ctx).
		Table(s.tableName).
		Where("id IN ?", ids).
		Update("published_at", time.Now())

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (o *outboxMessage) ID() string          { return o.id }
func (o *outboxMessage) MessageName() string { return o.name }
func (o *outboxMessage) Subject() string     { return o.subject }
func (o *outboxMessage) Data() []byte        { return o.data }
