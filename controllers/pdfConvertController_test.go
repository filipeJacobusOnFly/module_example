package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPdfHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/pdf", PdfHandler)

	t.Run("should return 400 when URL is not provided", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/pdf", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "URL não fornecida")
	})

	t.Run("should return 200 and PDF content when URL is provided", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/pdf?url=http://example.com", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "application/pdf", w.Header().Get("Content-Type"))
		assert.Equal(t, "attachment; filename=output.pdf", w.Header().Get("Content-Disposition"))

		assert.NotEmpty(t, w.Body.Bytes())
	})

}
