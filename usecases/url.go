package usecases

import (
	"context"
	"fmt"
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

func (u urlUsecase) GenerateCode(ctx context.Context) (string, utils.AppErr) {
	code := utils.GetShortKey(5)

	_, err := u.repo.GetByCode(ctx, code)
	if err != nil {
		if err.ErrCode() == utils.ERR_OBJ_NOT_FOUND {
			return code, nil
		} else {
			return "", err
		}
	} else {
		fmt.Println(3)
		return u.GenerateCode(ctx)
	}
}

func (u urlUsecase) Create(ctx context.Context, url string) (domain.URLCreateResp, utils.AppErr) {
	res := domain.URLCreateResp{}

	code, err := u.GenerateCode(ctx)
	if err != nil {
		return res, err
	}

	dbRes, err := u.repo.Create(ctx, domain.URLDB{Code: code, URL: url, CreatedAt: time.Now()})
	if err != nil {
		return res, err
	}

	res, err = dbRes.ToCreateRes()
	if err != nil {
		return res, err
	}
	return res, nil
}

func (u urlUsecase) Get(ctx context.Context, code string) (string, utils.AppErr) {
	dbRes, err := u.repo.GetByCode(ctx, code)
	if err != nil {
		return "", err
	}

	return dbRes.URL, err
}

func (u urlUsecase) List(ctx context.Context, req domain.URLListReq) (domain.URLListResp, utils.AppErr) {
	res := domain.URLListResp{}

	count, dbRes, err := u.repo.List(ctx, req)
	if err != nil {
		return res, err
	}

	res.Count = count
	res.Results = dbRes
	return res, nil
}
