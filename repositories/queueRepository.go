package repositories

import (
	"module_example/models"
)

var RecordChannel = make(chan models.Record, 10000)

type RecordRepositoryInterface interface {
	CreateRecords(records []models.Record) error
}

func StartBatchProcessing(repo RecordRepositoryInterface) {
	var records []models.Record

	for record := range RecordChannel {
		records = append(records, record)

		if len(records) >= 5000 {
			if err := repo.CreateRecords(records); err != nil {
				continue
			}
			records = nil
		}
	}
}
