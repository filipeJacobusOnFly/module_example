package database

import (
	"database/sql"
	"module_example/src/http/models"

	"github.com/sirupsen/logrus"
)

type RecordRepository struct {
	DB *sql.DB
}

func NewRecordRepository(db *sql.DB) *RecordRepository {
	return &RecordRepository{DB: db}
}

func (r *RecordRepository) CreateRecords(records []models.Record) error {
	tx, err := r.DB.Begin()
	if err != nil {
		logrus.Errorf("Erro ao iniciar a transação: %v", err)
		return err
	}
	defer func() {
		if err != nil {
			logrus.Warn("Rollback da transação devido a erro")
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				logrus.Errorf("Erro ao fazer rollback: %v", rollbackErr)
			}
		} else {
			logrus.Info("Transação confirmada com sucesso")
			err = tx.Commit()
		}
	}()

	stmt, err := tx.Prepare("INSERT INTO records (record_id, date) VALUES ($1, $2)")
	if err != nil {
		logrus.Errorf("Erro ao preparar a instrução SQL: %v", err)
		return err
	}
	defer stmt.Close()

	for _, record := range records {
		_, err := stmt.Exec(record.RecordID, record.Date)
		if err != nil {
			logrus.Errorf("Erro ao inserir registro: %+v, erro: %v", record, err)
			return err
		}
		logrus.Infof("Registro inserido com sucesso: %+v", record)
	}

	return nil
}
