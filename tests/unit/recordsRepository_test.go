package unit

import (
	"database/sql"
	"testing"
	"time"

	"module_example/src/http/models"
	repositories "module_example/src/http/repository"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func TestCreateRecords(t *testing.T) {

	logrus.SetLevel(logrus.InfoLevel)

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		logrus.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE records (record_id INTEGER PRIMARY KEY, date TEXT)")
	if err != nil {
		logrus.Fatalf("failed to create table: %v", err)
	}

	repo := repositories.NewRecordRepository(db)

	records := []models.Record{
		{RecordID: 1, Date: time.Now()},
		{RecordID: 2, Date: time.Now()},
	}

	logrus.Infof("Creating records: %+v", records)
	err = repo.CreateRecords(records)
	if err != nil {
		logrus.Fatalf("CreateRecords failed: %v", err)
	}

	rows, err := db.Query("SELECT record_id, date FROM records")
	if err != nil {
		logrus.Fatalf("failed to query records: %v", err)
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		count++
	}

	logrus.Infof("Queried %d records from the database", count)

	if count != len(records) {
		t.Errorf("expected %d records, got %d", len(records), count)
		logrus.Errorf("Record count mismatch: expected %d, got %d", len(records), count)
	} else {
		logrus.Info("Record creation test passed successfully")
	}
}
