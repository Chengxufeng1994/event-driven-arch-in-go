package model

import "time"

type StoreCache struct {
	ID        string    `gorm:"primaryKey;not null"`
	Name      string    `gorm:"column:name;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (StoreCache) TableName() string {
	return "search.stores_cache"
}
