package repositories

import (
	"database/sql"
	"module_example/cache"
	"module_example/structs"
)

type TokenRepository struct {
	DB    *sql.DB
	Cache *cache.TokenCache
}

func NewTokenRepository(db *sql.DB, cache *cache.TokenCache) *TokenRepository {
	return &TokenRepository{DB: db, Cache: cache}
}

func (r *TokenRepository) GetToken(tokenValue string) (*structs.Token, error) {
	if token, exists := r.Cache.GetToken(tokenValue); exists {
		return token, nil
	}

	var token structs.Token

	query := "SELECT id, token, exp_date FROM tokens WHERE token = $1"
	row := r.DB.QueryRow(query, tokenValue)

	err := row.Scan(&token.ID, &token.Token, &token.ExpDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	r.Cache.SetToken(tokenValue, &token)

	return &token, nil
}

func (r *TokenRepository) CreateToken(token structs.Token) error {
	query := "INSERT INTO tokens (token, exp_date) VALUES ($1, $2)"
	_, err := r.DB.Exec(query, token.Token, token.ExpDate)
	return err
}
