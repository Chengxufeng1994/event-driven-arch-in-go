package model

type Event struct {
	StreamID      string `gorm:"primary_key;column:stream_id;not nul"`
	StreamName    string `gorm:"primary_key;column:stream_name;not nul"`
	StreamVersion int    `gorm:"primary_key;column:stream_version;not nul"`
	EventID       string `gorm:"column:event_id;not nul"`
	EventName     string `gorm:"column:event_name;not nul"`
	EventData     []byte `gorm:"column:event_data;not nul"`
	OccurredAt    int64  `gorm:"column:occurred_at;default:CURRENT_TIMESTAMP"`
}
