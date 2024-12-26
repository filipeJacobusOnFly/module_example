package controllers

import (
	"net/http"

	"module_example/src/http/models"
	repositories "module_example/src/http/repository"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RecordHandler(repo *repositories.RecordRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var record models.Record

		if err := c.ShouldBindJSON(&record); err != nil {
			logrus.WithError(err).Warn("Dados inválidos recebidos")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
			return
		}

		if record.RecordID == 0 || record.Date.IsZero() {
			logrus.Warn("Dados inválidos: RecordID ou Date não fornecidos")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
			return
		}

		logrus.Infof("Registro recebido: %+v", record)
		c.JSON(http.StatusAccepted, gin.H{"message": "Registro recebido"})
	}
}
