package unit

import (
	"module_example/src/http/models"
	repositories "module_example/src/http/repository"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

type MockRecordRepository struct {
	Records []models.Record
	Err     error
}

func (m *MockRecordRepository) CreateRecords(records []models.Record) error {
	if m.Err != nil {
		return m.Err
	}
	m.Records = append(m.Records, records...)
	return nil
}

func TestStartBatchProcessing(t *testing.T) {

	logrus.SetLevel(logrus.InfoLevel)

	repo := &MockRecordRepository{}
	logrus.Info("Starting batch processing")
	go repositories.StartBatchProcessing(repo)

	for i := 0; i < 10000; i++ {
		repositories.RecordChannel <- models.Record{
			RecordID: uint(i),
			Date:     time.Now(),
		}
	}

	close(repositories.RecordChannel)

	time.Sleep(1 * time.Second)

	logrus.Infof("Expected to have processed 10000 records, currently have %d", len(repo.Records))

	if len(repo.Records) != 10000 {
		t.Errorf("Esperado 10000 registros, mas obteve %d", len(repo.Records))
		logrus.Errorf("Test failed: expected 10000 records, but got %d", len(repo.Records))
	}

	for i, record := range repo.Records {
		if record.RecordID != uint(i) {
			t.Errorf("Registro %d tem RecordID %d, esperado %d", i, record.RecordID, i)
			logrus.Errorf("Record mismatch: expected RecordID %d, but got %d", i, record.RecordID)
		}
	}

	logrus.Info("Batch processing test completed successfully")
}
