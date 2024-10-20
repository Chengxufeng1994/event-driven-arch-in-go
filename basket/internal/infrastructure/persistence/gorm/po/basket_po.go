package po

import (
	"time"
)

type Basket struct {
	ID         string    `gorm:"primaryKey;not null"`
	CustomerID string    `gorm:"not null"`
	PaymentID  string    `gorm:"not null"`
	Items      []byte    `gorm:"column:items"`
	Status     string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

func (b *Basket) TableName() string {
	return "baskets.baskets"
}
