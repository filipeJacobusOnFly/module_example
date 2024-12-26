package controllers

import (
	"fmt"
	"net/http"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func convertHtmlToPdf(url string) ([]byte, error) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		logrus.WithError(err).Error("Erro ao criar gerador de PDF")
		return nil, err
	}

	page := wkhtmltopdf.NewPage(url)
	pdfg.AddPage(page)

	err = pdfg.Create()
	if err != nil {
		logrus.WithError(err).Error("Erro ao criar PDF")
		return nil, err
	}

	logrus.Infof("PDF gerado com sucesso para a URL: %s", url)
	return pdfg.Bytes(), nil
}

func PdfHandler(c *gin.Context) {
	url := c.Query("url")

	if url == "" {
		logrus.Warn("URL não fornecida")
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL não fornecida"})
		return
	}

	pdfContent, err := convertHtmlToPdf(url)
	if err != nil {
		logrus.WithError(err).Error("Erro ao gerar PDF")
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao gerar PDF: %v", err)})
		return
	}

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=output.pdf")
	c.Data(http.StatusOK, "application/pdf", pdfContent)
}
