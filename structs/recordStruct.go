package structs

import "time"

type Record struct {
	RecordID uint      `gorm:"primaryKey"`
	Date     time.Time `gorm:"not null"`
}
