package repositories

import (
	"database/sql"
	"module_example/structs"
)

type RecordRepository struct {
	DB *sql.DB
}

func NewRecordRepository(db *sql.DB) *RecordRepository {
	return &RecordRepository{DB: db}
}

func (r *RecordRepository) CreateRecords(records []structs.Record) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	stmt, err := tx.Prepare("INSERT INTO records (record_id, date) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, record := range records {
		_, err := stmt.Exec(record.RecordID, record.Date)
		if err != nil {
			return err
		}
	}

	return nil
}
