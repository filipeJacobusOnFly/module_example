package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"module_example/controllers"
	"module_example/models"
	"module_example/repositories"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRecordHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := &repositories.RecordRepository{}
	repositories.RecordChannel = make(chan models.Record, 10000)

	defer close(repositories.RecordChannel)

	router := gin.Default()
	router.POST("/record", controllers.RecordHandler(repo))

	t.Run("Deve retornar 202 ao receber um registro válido", func(t *testing.T) {
		record := models.Record{
			RecordID: 1,
			Date:     time.Now(),
		}
		jsonRecord, err := json.Marshal(record)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/record", bytes.NewBuffer(jsonRecord))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusAccepted, resp.Code)

		var response map[string]string
		err = json.Unmarshal(resp.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Registro recebido", response["message"])
	})

	t.Run("Deve retornar 400 ao receber dados inválidos", func(t *testing.T) {
		invalidJSON := []byte(`{"invalid_field": "value"}`)

		req, err := http.NewRequest(http.MethodPost, "/record", bytes.NewBuffer(invalidJSON))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)

		var response map[string]string
		err = json.Unmarshal(resp.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Dados inválidos", response["error"])
	})
}
