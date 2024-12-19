package controllers

import (
	"module_example/structs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTokenRepository struct {
	mock.Mock
}

func (m *MockTokenRepository) GetToken(token string) (*structs.Token, error) {
	args := m.Called(token)
	return args.Get(0).(*structs.Token), args.Error(1)
}

func TestAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name               string
		authHeader         string
		expectedStatusCode int
		expectedResponse   string
		mockToken          *structs.Token
		mockError          error
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockRepo := new(MockTokenRepository)
			if tt.mockToken != nil {
				mockRepo.On("GetToken", "valid_token").Return(tt.mockToken, tt.mockError)
			} else {
				mockRepo.On("GetToken", "expired_token").Return(tt.mockToken, tt.mockError)
			}

			r := gin.New()
			r.Use(AuthMiddleware(mockRepo))
			r.GET("/test", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "success"})
			})

			req, _ := http.NewRequest(http.MethodGet, "/test", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			if tt.expectedResponse != "" {
				assert.JSONEq(t, tt.expectedResponse, w.Body.String())
			}
		})
	}
}
