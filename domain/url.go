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

type ListShortURLReq struct {
	Page  int64
	Limit int64
}

type ListShortURLRes struct {
	Count   int
	Results []ShortURL
}

type URLUsecase interface {
	Create(ctx context.Context, url string) (string, utils.AppErr)
	Get(ctx context.Context, code string) (string, utils.AppErr)
	List(ctx context.Context, req ListShortURLReq) (ListShortURLRes, utils.AppErr)
}

type URLRepository interface {
	Create(ctx context.Context, req ShortURL) (ShortURL, utils.AppErr)
	GetByURL(ctx context.Context, url string) (ShortURL, utils.AppErr)
	GetByCode(ctx context.Context, code string) (ShortURL, utils.AppErr)
	List(ctx context.Context, req ListShortURLReq) (int, []ShortURL, utils.AppErr)
}
