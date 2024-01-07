package domain

import (
	"context"
	"net/url"
	"os"
	"shorturl/utils"
	"time"
)

type URLDB struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Code      string    `json:"code"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

type URLCreateReq struct {
	URL string `json:"url" validate:"required,url"`
}

type URLCreateResp struct {
	ShortURL string `json:"short_url"`
}

type URLListReq struct {
	Page  int64
	Limit int64
}

type URLListResp struct {
	Count   int     `json:"count"`
	Results []URLDB `json:"result"`
}

type URLUsecase interface {
	Create(ctx context.Context, url string, userID uint) (URLCreateResp, utils.AppErr)
	Get(ctx context.Context, code string) (string, utils.AppErr)
	List(ctx context.Context, req URLListReq) (URLListResp, utils.AppErr)
}

type URLRepository interface {
	Create(ctx context.Context, req URLDB) (URLDB, utils.AppErr)
	GetByURL(ctx context.Context, url string) (URLDB, utils.AppErr)
	GetByCode(ctx context.Context, code string) (URLDB, utils.AppErr)
	List(ctx context.Context, req URLListReq) (int, []URLDB, utils.AppErr)
}

func (r URLDB) ToCreateRes() (URLCreateResp, utils.AppErr) {
	fullPath, err := url.JoinPath(os.Getenv("BASE_URL"), r.Code)
	if err != nil {
		return URLCreateResp{}, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
	}
	return URLCreateResp{ShortURL: fullPath}, nil
}
