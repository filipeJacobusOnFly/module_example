package controllers

import (
	"net/http"

	"module_example/repositories"
	"module_example/structs"

	"github.com/gin-gonic/gin"
)

func RecordHandler(repo *repositories.RecordRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var record structs.Record

		if err := c.ShouldBindJSON(&record); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
			return
		}

		if record.RecordID == 0 || record.Date.IsZero() {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
			return
		}

		repositories.RecordChannel <- record

		c.JSON(http.StatusAccepted, gin.H{"message": "Registro recebido"})
	}
}
