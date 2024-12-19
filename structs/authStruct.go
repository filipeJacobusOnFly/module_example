package structs

import (
	"time"
)

type Token struct {
	ID      uint      `json:"id"`
	Token   string    `json:"token"`
	ExpDate time.Time `json:"date"`
}
