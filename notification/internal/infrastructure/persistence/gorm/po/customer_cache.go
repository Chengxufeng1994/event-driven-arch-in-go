package po

import "time"

type CustomerCache struct {
	ID        string    `gorm:"primaryKey;not null"`
	Name      string    `gorm:"column:name;not null"`
	SmsNumber string    `gorm:"column:sms_number;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (c *CustomerCache) TableName() string {
	return "notifications.customers_cache"
}
