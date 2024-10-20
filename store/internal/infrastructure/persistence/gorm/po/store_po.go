package po

import "time"

type Store struct {
	ID            string    `gorm:"primaryKey;not null"`
	Name          string    `gorm:"not null"`
	Location      string    `gorm:"not null"`
	Participating bool      `gorm:"index;not null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}

func (store *Store) TableName() string {
	return "stores.stores"
}
