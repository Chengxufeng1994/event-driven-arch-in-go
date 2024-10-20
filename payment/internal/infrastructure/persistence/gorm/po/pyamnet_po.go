package po

import (
	"time"

	"github.com/shopspring/decimal"
)

type Payment struct {
	ID         string          `gorm:"primaryKey;not null"`
	CustomerID string          `gorm:"not null;index"`
	Amount     decimal.Decimal `gorm:"type:decimal(9,4);not null"`
	CreatedAt  time.Time       `gorm:"autoCreateTime"`
	UpdatedAt  time.Time       `gorm:"autoUpdateTime"`
}

func (Payment) TableName() string {
	return "payments.payments"
}
