package binlogger

type DbEvent struct {
	Id        int    `gorm:"column:id"`
	EventType string `gorm:"column:event_type"`
	RecordId  int    `gorm:"column:record_id"`
	Table     string `gorm:"column:event_table"`
}

// func (User) TableName() string {
// 	return "user"
// }

// func (User) SchemaName() string {
// 	return "app_db"
// }
