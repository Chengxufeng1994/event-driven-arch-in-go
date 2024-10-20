package po

import "time"

type Customer struct {
	ID        string    `gorm:"primaryKey;not null"`
	Name      string    `gorm:"not null"`
	SmsNumber string    `gorm:"not null"`
	Enabled   bool      `gorm:"not null;default:true"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (c *Customer) TableName() string {
	return "customers.customers"
}
