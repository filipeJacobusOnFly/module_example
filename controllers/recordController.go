package controllers

import (
	"net/http"
	"time"

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

		record.Date = time.Now()

		repositories.RecordChannel <- record

		c.JSON(http.StatusAccepted, gin.H{"message": "Registro recebido"})
	}
}
