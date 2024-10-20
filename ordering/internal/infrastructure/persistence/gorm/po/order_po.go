package po

import (
	"time"
)

type Order struct {
	ID         string    `gorm:"primaryKey;not null"`
	CustomerID string    `gorm:"column:customer_id;not null"`
	PaymentID  string    `gorm:"column:payment_id;not null"`
	InvoiceID  string    `gorm:"column:invoice_id;not null"`
	ShoppingID string    `gorm:"column:shopping_id;not null"`
	Items      []byte    `gorm:"column:items"`
	Status     string    `gorm:"column:status;not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

func (o *Order) TableName() string {
	return "ordering.orders"
}
