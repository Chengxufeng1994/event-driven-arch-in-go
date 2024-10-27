package po

import (
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID          string          `gorm:"primaryKey;not null"`
	StoreID     string          `gorm:"not null;index"`
	Name        string          `gorm:"not null"`
	Description string          `gorm:"not null"`
	SKU         string          `gorm:"not null"`
	Price       decimal.Decimal `gorm:"type:decimal(9,4);not null"`
	CreatedAt   time.Time       `gorm:"autoCreateTime"`
	UpdatedAt   time.Time       `gorm:"autoUpdateTime"`
}

func (*Product) TableName() string {
	return "stores.products"
}
