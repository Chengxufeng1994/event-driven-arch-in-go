package po

import (
	"time"

	"github.com/shopspring/decimal"
)

type ProductCache struct {
	ID        string          `gorm:"primaryKey;not null"`
	StoreID   string          `gorm:"not null"`
	Name      string          `gorm:"not null"`
	Price     decimal.Decimal `gorm:"type:decimal(9,4);not null"`
	CreatedAt time.Time       `gorm:"autoCreateTime"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime"`
}

func (pc ProductCache) TableName() string {
	return "baskets.products_cache"
}
