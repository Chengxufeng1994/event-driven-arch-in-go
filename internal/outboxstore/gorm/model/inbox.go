package model

import "time"

type Inbox struct {
	ID         string    `gorm:"primaryKey;not null"`
	Name       string    `gorm:"not null"`
	Subject    string    `gorm:"not null"`
	Data       []byte    `gorm:"not null"`
	Metadata   []byte    `gorm."not null"`
	SentAt     time.Time `gorm:"type:timestamptz"`
	ReceivedAt time.Time `gorm:"type:timestamptz;not null"`
}
