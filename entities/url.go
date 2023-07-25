package entities

import (
	"context"
	"time"
)

type ShortURL struct {
	ID        int64     `json:"id"`
	Code      string    `json:"code"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

type URLUsecase interface {
	Create(ctx context.Context, url string) (string, error)
	Get(ctx context.Context, code string) (string, error)
}

type URLRepository interface {
	Create(ctx context.Context, req ShortURL) (ShortURL, error)
	GetByURL(ctx context.Context, url string) (ShortURL, error)
	GetByCode(ctx context.Context, code string) (ShortURL, error)
	List(ctx context.Context, page int64, limit int64) ([]ShortURL, error)
}
