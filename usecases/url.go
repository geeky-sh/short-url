package usecases

import (
	"context"
	"shorturl/domain"
	"shorturl/utils"
	"time"
)

type urlUsecase struct {
	repo domain.URLRepository
}

func NewURLUsecase(r domain.URLRepository) domain.URLUsecase {
	return urlUsecase{repo: r}
}

func (u urlUsecase) Create(ctx context.Context, url string) (string, utils.AppErr) {
	code := utils.GetUniq(url)

	dbRes, err := u.repo.GetByCode(ctx, code)
	if err != nil && err.ErrCode() == utils.ERR_OBJ_NOT_FOUND {
		_, err := u.repo.Create(ctx, domain.ShortURL{Code: code, URL: url, CreatedAt: time.Now()})
		if err != nil {
			return "", err
		}

		return code, nil
	} else if err != nil {
		return "", err
	} else {
		return dbRes.Code, nil
	}
}

func (u urlUsecase) Get(ctx context.Context, code string) (string, utils.AppErr) {
	dbRes, err := u.repo.GetByCode(ctx, code)
	if err != nil {
		return "", err
	}

	return dbRes.URL, err
}

func (u urlUsecase) List(ctx context.Context, req domain.ListShortURLReq) (domain.ListShortURLRes, utils.AppErr) {
	res := domain.ListShortURLRes{}

	count, dbRes, err := u.repo.List(ctx, req)
	if err != nil {
		return res, err
	}

	res.Count = count
	res.Results = dbRes
	return res, nil
}
