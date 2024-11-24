package gorm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
	"gorm.io/gorm"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/outboxstore/gorm/model"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/tm"
)

type OutboxStore struct {
	tableName string
	db        *gorm.DB
}

type outboxMessage struct {
	id       string
	name     string
	subject  string
	data     []byte
	metadata ddd.Metadata
	sentAt   time.Time
}

var _ tm.OutboxStore = (*OutboxStore)(nil)
var _ am.Message = (*outboxMessage)(nil)

func NewOutboxStore(tableName string, db *gorm.DB) *OutboxStore {
	return &OutboxStore{
		tableName: tableName,
		db:        db,
	}
}

// Save implements tm.OutboxStore.
func (s *OutboxStore) Save(ctx context.Context, msg am.Message) error {
	const query = "INSERT INTO %s (id, name, subject, data, metadata, sent_at) VALUES ($1, $2, $3, $4, $5, $6)"

	metadata, err := json.Marshal(msg.Metadata())
	if err != nil {
		return err
	}

	stmt := fmt.Sprintf(query, s.tableName)
	result := s.db.WithContext(ctx).
		Exec(stmt, msg.ID(), msg.MessageName(), msg.Subject(), msg.Data(), metadata, msg.SentAt())

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
func (s *OutboxStore) FindUnpublished(ctx context.Context, limit int) ([]am.Message, error) {
	var rows []model.Outbox

	err := s.db.WithContext(ctx).
		Table(s.tableName).
		Where("published_at IS NULL").
		Limit(limit).
		Find(&rows).
		Error
	if err != nil {
		return nil, err
	}

	outboxMsgs := make([]am.Message, 0, len(rows))
	for _, row := range rows {
		msg := outboxMessage{
			id:      row.ID,
			name:    row.Name,
			subject: row.Subject,
			data:    row.Data,
			sentAt:  row.SentAt,
		}

		err = json.Unmarshal(row.Metadata, &msg.metadata)
		if err != nil {
			return nil, err
		}

		outboxMsgs = append(outboxMsgs, msg)
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

func (m outboxMessage) ID() string             { return m.id }
func (m outboxMessage) Subject() string        { return m.subject }
func (m outboxMessage) MessageName() string    { return m.name }
func (m outboxMessage) Data() []byte           { return m.data }
func (m outboxMessage) Metadata() ddd.Metadata { return m.metadata }
func (m outboxMessage) SentAt() time.Time      { return m.sentAt }
