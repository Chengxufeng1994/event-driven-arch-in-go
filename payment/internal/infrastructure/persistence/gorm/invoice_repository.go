package gorm

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/infrastructure/persistence/gorm/mapper"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/internal/infrastructure/persistence/gorm/po"
	"github.com/stackus/errors"
)

type GormInvoiceRepository struct {
	db     *gorm.DB
	mapper mapper.InvoiceMapperIntf
}

var _ repository.InvoiceRepository = (*GormInvoiceRepository)(nil)

func NewGormInvoiceRepository(db *gorm.DB) *GormInvoiceRepository {
	return &GormInvoiceRepository{
		db:     db,
		mapper: mapper.NewInvoiceMapper(),
	}
}

// Save implements repository.InvoiceRepository.
func (r *GormInvoiceRepository) Save(ctx context.Context, invoice *aggregate.InvoiceAgg) error {
	invoicePO, err := r.mapper.ToPersistence(invoice)
	if err != nil {
		return err
	}

	tx := r.db.Begin()
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.WithContext(ctx).
		Model(&po.Invoice{}).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"order_id", "status", "amount"}),
		}).
		Create(&invoicePO).Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, "inserting invoice")
	}

	return tx.Commit().Error
}

// Update implements repository.InvoiceRepository.
func (r *GormInvoiceRepository) Update(ctx context.Context, invoice *aggregate.InvoiceAgg) error {
	invoicePO, err := r.mapper.ToPersistence(invoice)
	if err != nil {
		return err
	}

	tx := r.db.Begin()
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.WithContext(ctx).
		Model(&po.Invoice{}).
		Where("id = ?", invoicePO.ID).
		Select("order_id", "status", "amount").
		Updates(&po.Invoice{
			OrderID: invoicePO.OrderID,
			Status:  invoicePO.Status,
			Amount:  invoicePO.Amount,
		}).
		Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, "updating invoice")
	}

	return tx.Commit().Error
}

// Find implements repository.InvoiceRepository.
func (r *GormInvoiceRepository) Find(ctx context.Context, invoiceID string) (*aggregate.InvoiceAgg, error) {
	var invoice *po.Invoice

	result := r.db.WithContext(ctx).
		Model(&po.Invoice{}).
		Where("id = ?", invoiceID).
		First(&invoice)

	err := result.Error
	if err != nil {
		// if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// 	return nil, nil
		// }

		return nil, errors.Wrap(err, "finding invoice")
	}

	return r.mapper.ToDomain(invoice)
}
