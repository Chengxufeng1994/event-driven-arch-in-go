package model

import "time"

type CustomerCache struct {
	ID        string    `gorm:"primaryKey;not null"`
	Name      string    `gorm:"column:name;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (CustomerCache) TableName() string {
	return "search.customers_cache"
}
