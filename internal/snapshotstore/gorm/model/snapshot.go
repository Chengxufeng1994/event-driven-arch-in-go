package model

type Snapshot struct {
	StreamID      string `gorm:"primary_key;column:stream_id;not nul"`
	StreamName    string `gorm:"primary_key;column:stream_name;not nul"`
	StreamVersion int    `gorm:"column:stream_version;not nul"`
	SnapshotName  string `gorm:"column:snapshot_name;not nul"`
	SnapshotData  []byte `gorm:"column:snapshot_data;not nul"`
	UpdatedAt     int64  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
}
