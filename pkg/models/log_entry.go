package models

import (
	"time"
)

type LogEntry struct {
	ID        int       `gorm:"primaryKey"`
	IPAddress string    `gorm:"column:ip_address"`
	Timestamp time.Time `gorm:"column:timestamp"`
	URL       string    `gorm:"column:url"`
	Method    string    `gorm:"column:method"`
	Status    int       `gorm:"column:status"`
	Latency   float64   `gorm:"column:latency"`
}
