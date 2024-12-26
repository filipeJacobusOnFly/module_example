package controllers

import (
	"net/http"

	"module_example/src/http/models"
	repositories "module_example/src/http/repository"

	"github.com/gin-gonic/gin"
)

func RecordHandler(repo *repositories.RecordRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var record models.Record

		if err := c.ShouldBindJSON(&record); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
			return
		}

		if record.RecordID == 0 || record.Date.IsZero() {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
			return
		}

		c.JSON(http.StatusAccepted, gin.H{"message": "Registro recebido"})
	}
}
