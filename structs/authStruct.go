package structs

import (
	"time"
)

type Token struct {
	ID      uint      `gorm:"primaryKey"`
	Token   string    `gorm:"uniqueIndex"`
	ExpDate time.Time `gorm:"not null"`
}
