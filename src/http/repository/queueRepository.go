package database

import (
	"module_example/src/http/models"

	"github.com/sirupsen/logrus"
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
			logrus.Infof("Processando %d registros em lote", len(records))
			if err := repo.CreateRecords(records); err != nil {
				logrus.Errorf("Erro ao criar registros em lote: %v", err)
				// Você pode optar por continuar ou lidar com o erro de outra forma
				continue
			}
			logrus.Infof("Registros em lote criados com sucesso: %d", len(records))
			records = nil
		}
	}

	// Se houver registros restantes que não foram processados
	if len(records) > 0 {
		logrus.Infof("Processando %d registros restantes em lote", len(records))
		if err := repo.CreateRecords(records); err != nil {
			logrus.Errorf("Erro ao criar registros restantes em lote: %v", err)
		} else {
			logrus.Infof("Registros restantes criados com sucesso: %d", len(records))
		}
	}
}
