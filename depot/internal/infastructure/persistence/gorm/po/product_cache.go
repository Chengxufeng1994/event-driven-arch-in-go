package po

import "time"

type ProductCache struct {
	ID        string    `gorm:"primaryKey;not null"`
	StoreID   string    `gorm:"not null"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (pc ProductCache) TableName() string {
	return "depot.products_cache"
}
