package repositories

import (
	"database/sql"
	"module_example/cache"
	"module_example/models"
)

type DBInterface interface {
	QueryRow(query string, args ...interface{}) RowInterface
	Exec(query string, args ...interface{}) (sql.Result, error)
}
type DBWrapper struct {
	*sql.DB
}

func (db *DBWrapper) QueryRow(query string, args ...interface{}) RowInterface {
	return db.DB.QueryRow(query, args...)
}

func (db *DBWrapper) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.DB.Exec(query, args...)
}

type TokenRepositoryInterface interface {
	GetToken(token string) (*models.Token, error)
}
type RowInterface interface {
	Scan(dest ...interface{}) error
}
type TokenRepository struct {
	DB    DBInterface
	Cache cache.TokenCacheInterface
}

func NewTokenRepository(db DBInterface, cache cache.TokenCacheInterface) *TokenRepository {
	return &TokenRepository{DB: db, Cache: cache}
}
func (r *TokenRepository) GetToken(tokenValue string) (*models.Token, error) {
	if token, exists := r.Cache.GetToken(tokenValue); exists {
		return token, nil
	}

	var token models.Token

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

func (r *TokenRepository) CreateToken(token models.Token) error {
	query := "INSERT INTO tokens (token, exp_date) VALUES ($1, $2)"
	_, err := r.DB.Exec(query, token.Token, token.ExpDate)
	return err
}
