package model

import "time"

type Outbox struct {
	ID          string     `gorm:"primaryKey;not null"`
	Name        string     `gorm:"not null"`
	Subject     string     `gorm:"not null"`
	Data        []byte     `gorm:"not null"`
	Metadata    []byte     `gorm."not null"`
	SentAt      time.Time  `gorm:"type:timestamptz"`
	PublishedAt *time.Time `gorm:"type:timestamptz"`
}
