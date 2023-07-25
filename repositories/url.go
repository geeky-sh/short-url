package repositories

import (
	"context"
	"shorturl/entities"

	"github.com/jackc/pgx/v5"
)

const TABLE_NAME = "urls"

type urlRepository struct {
	db *pgx.Conn
	tb string
}

func NewURLRepository(db *pgx.Conn) entities.URLRepository {
	return urlRepository{db, TABLE_NAME}
}

func (r urlRepository) Create(ctx context.Context, req entities.ShortURL) (entities.ShortURL, error) {
	return entities.ShortURL{}, nil
}

func (r urlRepository) GetByURL(ctx context.Context, url string) (entities.ShortURL, error) {
	return entities.ShortURL{}, nil
}

func (r urlRepository) GetByCode(ctx context.Context, code string) (entities.ShortURL, error) {
	return entities.ShortURL{}, nil
}

func (r urlRepository) List(ctx context.Context, page int64, limit int64) ([]entities.ShortURL, error) {
	return []entities.ShortURL{}, nil
}
