package repositories

import (
	"module_example/structs"
	"testing"
	"time"
)

type MockRecordRepository struct {
	Records []structs.Record
	Err     error
}

func (m *MockRecordRepository) CreateRecords(records []structs.Record) error {
	if m.Err != nil {
		return m.Err
	}
	m.Records = append(m.Records, records...)
	return nil
}

func TestStartBatchProcessing(t *testing.T) {
	repo := &MockRecordRepository{}
	go StartBatchProcessing(repo)

	for i := 0; i < 10000; i++ {
		RecordChannel <- structs.Record{
			RecordID: uint(i),
			Date:     time.Now(),
		}
	}

	close(RecordChannel)

	time.Sleep(1 * time.Second)

	if len(repo.Records) != 10000 {
		t.Errorf("Esperado 10000 registros, mas obteve %d", len(repo.Records))
	}

	for i, record := range repo.Records {
		if record.RecordID != uint(i) {
			t.Errorf("Registro %d tem RecordID %d, esperado %d", i, record.RecordID, i)
		}
	}
}
