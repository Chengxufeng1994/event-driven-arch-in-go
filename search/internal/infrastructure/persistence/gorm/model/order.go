package model

import (
	"time"

	"github.com/lib/pq"
)

type Order struct {
	OrderID      string         `gorm:"primaryKey;not null"`
	CustomerID   string         `gorm:"column:customer_id;not null"`
	CustomerName string         `gorm:"column:customer_name;not null"`
	Items        []byte         `gorm:"column:items;not null"`
	Status       string         `gorm:"column:status;not null"`
	ProductIDs   pq.StringArray `gorm:"column:product_ids;type:text[];not null"`
	StoreIDs     pq.StringArray `gorm:"column:store_ids;type:text[];not null"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
}

func (*Order) TableName() string {
	return "search.orders"
}
