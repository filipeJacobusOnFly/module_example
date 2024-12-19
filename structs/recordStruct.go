package structs

import "time"

type Record struct {
	RecordID uint      `json:"record_id"`
	Date     time.Time `json:"date"`
}
