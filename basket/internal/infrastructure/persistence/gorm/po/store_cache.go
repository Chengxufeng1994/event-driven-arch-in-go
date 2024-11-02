package po

import "time"

type StoreCache struct {
	ID        string    `gorm:"primaryKey;not null"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (sc StoreCache) TableName() string {
	return "baskets.stores_cache"
}
