package usecases

import (
	"context"
	"shorturl/entities"
)

type urlUsecase struct {
	repo entities.URLRepository
}

func NewURLUsecase(r entities.URLRepository) entities.URLUsecase {
	return urlUsecase{repo: r}
}

func (u urlUsecase) Create(ctx context.Context, url string) (string, error) {
	return "", nil
}

func (u urlUsecase) Get(ctx context.Context, code string) (string, error) {
	return "", nil
}
