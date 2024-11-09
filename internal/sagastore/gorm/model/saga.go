package model

import "time"

type Saga struct {
	ID           string    `gorm:"primaryKey;not null"`
	Name         string    `gorm:"not null"`
	Data         []byte    `gorm:"not null"`
	Step         int       `gorm:"not null"`
	Done         bool      `gorm:"not null"`
	Compensating bool      `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
