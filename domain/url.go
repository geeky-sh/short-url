package domain

import (
	"context"
	"shorturl/utils"
	"time"
)

type ShortURL struct {
	ID        uint      `json:"id"`
	Code      string    `json:"code"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateShortURL struct {
	URL string `json:"url" validate:"required,url"`
}

type ListShortURL struct {
	Page  int64
	Limit int64
}

type URLUsecase interface {
	Create(ctx context.Context, url string) (string, utils.AppErr)
	Get(ctx context.Context, code string) (string, utils.AppErr)
}

type URLRepository interface {
	Create(ctx context.Context, req ShortURL) (ShortURL, utils.AppErr)
	GetByURL(ctx context.Context, url string) (ShortURL, utils.AppErr)
	GetByCode(ctx context.Context, code string) (ShortURL, utils.AppErr)
	List(ctx context.Context, req ListShortURL) (int, []ShortURL, utils.AppErr)
}
