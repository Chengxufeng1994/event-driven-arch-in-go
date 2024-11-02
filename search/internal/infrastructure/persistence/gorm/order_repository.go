package gorm

import (
	"bytes"
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application/port/in/query"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/application/port/out"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/domain"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search/internal/infrastructure/persistence/gorm/model"
	"github.com/stackus/errors"
	"gorm.io/gorm"
)

type GormOrderRepository struct {
	db *gorm.DB
}

var _ out.OrderRepository = (*GormOrderRepository)(nil)

func NewGormOrderRepository(db *gorm.DB) *GormOrderRepository {
	return &GormOrderRepository{db}
}

// Add implements out.OrderRepository.
func (r *GormOrderRepository) Add(ctx context.Context, order *domain.Order) error {
	items, err := json.Marshal(order.Items)
	if err != nil {
		return err
	}

	storeMap := make(map[string]struct{})
	productIds := make([]string, len(order.Items))
	for i, item := range order.Items {
		productIds[i] = item.ProductID
		storeMap[item.StoreID] = struct{}{}
	}

	storeIds := make([]string, 0, len(storeMap))
	for storeId, _ := range storeMap {
		storeIds = append(storeIds, storeId)
	}

	result := r.db.WithContext(ctx).
		Model(&model.Order{}).
		Create(&model.Order{
			OrderID:      order.OrderID,
			CustomerID:   order.CustomerID,
			CustomerName: order.CustomerName,
			Items:        items,
			Status:       order.Status,
			ProductIDs:   productIds,
			StoreIDs:     storeIds,
			CreatedAt:    order.CreatedAt,
		})

	return result.Error
}

// UpdateStatus implements out.OrderRepository.
func (r *GormOrderRepository) UpdateStatus(ctx context.Context, orderID string, status string) error {
	result := r.db.WithContext(ctx).
		Model(&model.Order{}).
		Where("order_id = ?", orderID).
		Update("status", status)

	return result.Error
}

// Get implements out.OrderRepository.
func (r *GormOrderRepository) Get(ctx context.Context, orderID string) (*domain.Order, error) {
	var orderPO model.Order
	result := r.db.WithContext(ctx).
		Model(&model.Order{}).
		Where("order_id = ?", orderID).
		First(&orderPO)

	if result.Error != nil {
		return nil, result.Error
	}

	var items []domain.Item
	if err := json.Unmarshal(orderPO.Items, &items); err != nil {
		return nil, err
	}

	return &domain.Order{
		OrderID:      orderPO.OrderID,
		CustomerID:   orderPO.CustomerID,
		CustomerName: orderPO.CustomerName,
		Items:        items,
		Status:       orderPO.Status,
		CreatedAt:    orderPO.CreatedAt,
	}, nil
}

// Search implements out.OrderRepository.
func (r *GormOrderRepository) Search(ctx context.Context, search query.SearchOrders) ([]*domain.Order, error) {
	panic("unimplemented")
}

type IDArray []string

func (a *IDArray) Scan(src any) error {
	var sep = []byte(",")

	var data []byte
	switch v := src.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	default:
		return errors.ErrInvalidArgument.Msgf("IDArray: unsupported type: %T", src)
	}

	ids := make([]string, bytes.Count(data, sep))
	for i, id := range bytes.Split(bytes.Trim(data, "{}"), sep) {
		ids[i] = string(id)
	}

	*a = ids

	return nil
}

func (a IDArray) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}
	if len(a) == 0 {
		return "{}", nil
	}
	// unsafe way to do this; assumption is all ids are UUIDs
	return fmt.Sprintf("{%s}", strings.Join(a, ",")), nil
}
