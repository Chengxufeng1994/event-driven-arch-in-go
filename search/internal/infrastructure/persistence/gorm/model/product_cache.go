package model

import "time"

type ProductCache struct {
	ID        string    `gorm:"primaryKey;not null"`
	StoreID   string    `gorm:"column:store_id;not null"`
	Name      string    `gorm:"column:name;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (ProductCache) TableName() string {
	return "search.products_cache"
}
