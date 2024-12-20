package cache

import (
	"module_example/models"
	"sync"
)

type TokenCache struct {
	mu     sync.RWMutex
	tokens map[string]*models.Token
}

type TokenCacheInterface interface {
	GetToken(tokenValue string) (*models.Token, bool)
	SetToken(tokenValue string, token *models.Token)
}

func NewTokenCache() *TokenCache {
	return &TokenCache{
		tokens: make(map[string]*models.Token),
	}
}

func (c *TokenCache) GetToken(tokenValue string) (*models.Token, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	token, exists := c.tokens[tokenValue]
	return token, exists
}

func (c *TokenCache) SetToken(tokenValue string, token *models.Token) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.tokens[tokenValue] = token
}
