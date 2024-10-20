package po

import (
	"time"

	"github.com/shopspring/decimal"
)

type Invoice struct {
	ID        string          `gorm:"primaryKey;not null"`
	OrderID   string          `gorm:"not null;index"`
	Amount    decimal.Decimal `gorm:"type:decimal(9,4);not null"`
	Status    string          `gorm:"not null"`
	CreatedAt time.Time       `gorm:"autoCreateTime"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime"`
}

func (Invoice) TableName() string {
	return "payments.invoices"
}
