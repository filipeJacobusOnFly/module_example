package repositories

import (
	"module_example/structs"
)

var RecordChannel = make(chan structs.Record, 10000)

func StartBatchProcessing(repo *RecordRepository) {
	var records []structs.Record

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
