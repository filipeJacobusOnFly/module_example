package unit

import (
	controllers "module_example/src/http/controllers"
	"module_example/src/http/models"
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

func (m *MockTokenRepository) CreateToken(token models.Token) error {
	panic("unimplemented")
}

func (m *MockTokenRepository) GetToken(token string) (*models.Token, error) {
	args := m.Called(token)
	return args.Get(0).(*models.Token), args.Error(1)
}

func TestAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name               string
		authHeader         string
		expectedStatusCode int
		expectedResponse   string
		mockToken          *models.Token
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
			r.Use(controllers.AuthMiddleware(mockRepo))
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
