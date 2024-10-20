package po

import "time"

type ShoppingList struct {
	ID            string    `gorm:"primaryKey"`
	OrderID       string    `gorm:"column:order_id;not null"`
	Stops         []byte    `gorm:"column:stops"`
	AssignedBotID string    `gorm:"column:assigned_bot_id;not null"`
	Status        string    `gorm:"column:status;not null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}

func (ShoppingList) TableName() string {
	return "depot.shopping_lists"
}
