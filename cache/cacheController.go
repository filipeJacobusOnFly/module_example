package cache

import (
	"module_example/structs"
	"sync"
)

type TokenCache struct {
	mu     sync.RWMutex
	tokens map[string]*structs.Token
}

func NewTokenCache() *TokenCache {
	return &TokenCache{
		tokens: make(map[string]*structs.Token),
	}
}

func (c *TokenCache) GetToken(tokenValue string) (*structs.Token, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	token, exists := c.tokens[tokenValue]
	return token, exists
}

func (c *TokenCache) SetToken(tokenValue string, token *structs.Token) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.tokens[tokenValue] = token
}
