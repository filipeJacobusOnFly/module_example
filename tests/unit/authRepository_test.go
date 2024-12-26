package unit

import (
	"database/sql"
	"module_example/src/http/models"
	repositories "module_example/src/http/repository"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

type MockTokenCache struct {
	tokens map[string]*models.Token
}

func NewMockTokenCache() *MockTokenCache {
	return &MockTokenCache{tokens: make(map[string]*models.Token)}
}

func (m *MockTokenCache) GetToken(tokenValue string) (*models.Token, bool) {
	token, exists := m.tokens[tokenValue]
	return token, exists
}

func (m *MockTokenCache) SetToken(tokenValue string, token *models.Token) {
	m.tokens[tokenValue] = token
}

type TokenCache interface {
	GetToken(tokenValue string) (*models.Token, bool)
	SetToken(tokenValue string, token *models.Token)
}

type MockRow struct {
	token models.Token
	err   error
}

type MockDB struct {
	tokens map[string]models.Token
}

func (db *MockDB) QueryRow(query string, args ...interface{}) repositories.RowInterface {
	panic("unimplemented")
}

func NewMockDB() *MockDB {
	return &MockDB{tokens: make(map[string]models.Token)}
}

func (mr *MockRow) Scan(dest ...interface{}) error {
	if mr.err != nil {
		return mr.err
	}
	dest[0] = mr.token.ID
	dest[1] = mr.token.Token
	dest[2] = mr.token.ExpDate
	return nil
}

func (db *MockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	token := models.Token{
		Token:   args[0].(string),
		ExpDate: args[1].(time.Time),
	}
	db.tokens[token.Token] = token
	return nil, nil
}

func TestGetToken_CacheHit(t *testing.T) {
	logrus.Info("Iniciando o teste TestGetToken_CacheHit")

	cache := NewMockTokenCache()
	db := NewMockDB()
	repo := repositories.NewTokenRepository(db, cache)

	token := &models.Token{ID: 1, Token: "test-token", ExpDate: time.Now()}
	cache.SetToken("test-token", token)

	result, err := repo.GetToken("test-token")
	if err != nil {
		logrus.Errorf("Erro inesperado: %v", err)
		t.Fatalf("Erro inesperado: %v", err)
	}
	if result != token {
		logrus.Errorf("Esperado %v, mas obteve %v", token, result)
		t.Errorf("Esperado %v, mas obteve %v", token, result)
	} else {
		logrus.Infof("Teste TestGetToken_CacheHit passou com sucesso.")
	}
}

func TestCreateToken(t *testing.T) {
	logrus.Info("Iniciando o teste TestCreateToken")

	cache := NewMockTokenCache()
	db := NewMockDB()
	repo := repositories.NewTokenRepository(db, cache)

	token := models.Token{Token: "new-token", ExpDate: time.Now()}
	err := repo.CreateToken(token)
	if err != nil {
		logrus.Errorf("Erro inesperado: %v", err)
		t.Fatalf("Erro inesperado: %v", err)
	}

	if _, exists := db.tokens["new-token"]; !exists {
		logrus.Errorf("Esperado que o token 'new-token' existisse no banco de dados, mas não foi encontrado.")
		t.Errorf("Esperado que o token 'new-token' existisse no banco de dados, mas não foi encontrado.")
	} else {
		logrus.Infof("Token 'new-token' criado com sucesso.")
	}
}
