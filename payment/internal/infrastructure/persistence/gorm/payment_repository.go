package gorm

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/infrastructure/persistence/gorm/mapper"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/infrastructure/persistence/gorm/po"
	"github.com/stackus/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GormPaymentRepository struct {
	db     *gorm.DB
	mapper mapper.PaymentMapperIntf
}

var _ repository.PaymentRepository = (*GormPaymentRepository)(nil)

func NewGormPaymentRepository(db *gorm.DB) *GormPaymentRepository {
	return &GormPaymentRepository{
		db:     db,
		mapper: mapper.NewPaymentMapper(),
	}
}

func (r *GormPaymentRepository) Save(ctx context.Context, payment *aggregate.PaymentAgg) error {
	paymentPO, err := r.mapper.ToPersistent(payment)
	if err != nil {
		return errors.Wrap(err, "mapping payment")
	}

	result := r.db.WithContext(ctx).
		Model(&po.Payment{}).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"customer_id", "amount"}),
		}).
		Create(&paymentPO)

	return errors.Wrap(result.Error, "inserting payment")
}

func (r *GormPaymentRepository) Find(ctx context.Context, paymentID string) (*aggregate.PaymentAgg, error) {
	var payment *po.Payment
	result := r.db.WithContext(ctx).
		Where("id = ?", paymentID).
		First(&payment)

	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "finding payment")
	}

	return r.mapper.ToDomain(payment)
}
