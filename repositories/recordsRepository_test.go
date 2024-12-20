package repositories_test

import (
	"database/sql"
	"testing"
	"time"

	"module_example/models"
	"module_example/repositories"

	_ "github.com/mattn/go-sqlite3"
)

func TestCreateRecords(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE records (record_id INTEGER PRIMARY KEY, date TEXT)")
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}

	repo := repositories.NewRecordRepository(db)

	records := []models.Record{
		{RecordID: 1, Date: time.Now()},
		{RecordID: 2, Date: time.Now()},
	}

	err = repo.CreateRecords(records)
	if err != nil {
		t.Fatalf("CreateRecords failed: %v", err)
	}

	rows, err := db.Query("SELECT record_id, date FROM records")
	if err != nil {
		t.Fatalf("failed to query records: %v", err)
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		count++
	}

	if count != len(records) {
		t.Errorf("expected %d records, got %d", len(records), count)
	}
}
