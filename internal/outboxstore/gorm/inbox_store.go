package gorm

import (
	"context"
	"errors"
	"time"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/outboxstore/gorm/model"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/tm"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type InboxStore struct {
	tableName string
	db        *gorm.DB
}

var _ tm.InboxStore = (*InboxStore)(nil)

func NewInboxStore(tableName string, db *gorm.DB) *InboxStore {
	return &InboxStore{
		tableName: tableName,
		db:        db,
	}
}

func (s *InboxStore) Save(ctx context.Context, msg am.RawMessage) error {
	result := s.db.WithContext(ctx).
		Table(s.tableName).
		Create(&model.Inbox{
			ID:         msg.ID(),
			Name:       msg.MessageName(),
			Subject:    msg.Subject(),
			Data:       msg.Data(),
			ReceivedAt: time.Now(),
		})

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
