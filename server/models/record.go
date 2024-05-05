package models

import "time"

type Record struct {
	ID          int64      `gorm:"column:id;primaryKey"`
	ShortURLKey string     `gorm:"column:short_url_key"`
	LongURL     string     `gorm:"column:long_url"`
	Accesses    uint       `gorm:"column:accesses"`
	LastAccess  *time.Time `gorm:"column:last_access;type:timestamp"`
	CreatedAt   time.Time  `gorm:"column:created_at;type:timestamp"`
}

func (Record) TableName() string {
	return "records"
}
