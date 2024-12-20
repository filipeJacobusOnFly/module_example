package repositories

import (
	"database/sql"
	"testing"
	"time"

	"module_example/models"
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

func (db *MockDB) QueryRow(query string, args ...interface{}) RowInterface {
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
	cache := NewMockTokenCache()
	db := NewMockDB()
	repo := NewTokenRepository(db, cache)

	token := &models.Token{ID: 1, Token: "test-token", ExpDate: time.Now()}
	cache.SetToken("test-token", token)

	result, err := repo.GetToken("test-token")
	if err != nil {
		t.Fatalf("Erro inesperado: %v", err)
	}
	if result != token {
		t.Errorf("Esperado %v, mas obteve %v", token, result)
	}
}

func TestCreateToken(t *testing.T) {
	cache := NewMockTokenCache()
	db := NewMockDB()
	repo := NewTokenRepository(db, cache)

	token := models.Token{Token: "new-token", ExpDate: time.Now()}
	err := repo.CreateToken(token)
	if err != nil {
		t.Fatalf("Erro inesperado: %v", err)
	}

	if _, exists := db.tokens["new-token"]; !exists {
		t.Errorf("Esperado que o token 'new-token' existisse no banco de dados, mas não foi encontrado.")
	}
}
