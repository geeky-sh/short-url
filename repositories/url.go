package repositories

import (
	"context"
	"errors"
	"shorturl/domain"
	"shorturl/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const TABLE_NAME = "urls"

type urlRepository struct {
	db *pgxpool.Pool
	tb string
}

func NewURLRepository(db *pgxpool.Pool) domain.URLRepository {
	return urlRepository{db, TABLE_NAME}
}

func (r urlRepository) Create(ctx context.Context, req domain.ShortURL) (domain.ShortURL, utils.AppErr) {
	res := domain.ShortURL{}

	sql := `
	INSERT INTO short_urls (code, url, created_at)
	VALUES ($1, $2, $3) RETURNING id`
	_, err := r.db.Exec(ctx, sql, req.Code, req.URL, req.CreatedAt)
	if err != nil {
		return res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
	}

	return domain.ShortURL{}, nil
}

func (r urlRepository) GetByURL(ctx context.Context, url string) (domain.ShortURL, utils.AppErr) {
	res := domain.ShortURL{}

	sql := `
	SELECT id, code, url, created_at FROM short_urls
	WHERE url=$1`
	if err := r.db.QueryRow(ctx, sql, url).Scan(res.ID, res.Code, res.URL, res.CreatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return res, utils.NewAppErr(err.Error(), utils.ERR_OBJ_NOT_FOUND)
		}
		return res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
	}
	return res, nil
}

func (r urlRepository) GetByCode(ctx context.Context, code string) (domain.ShortURL, utils.AppErr) {
	res := domain.ShortURL{}

	sql := `
	SELECT id, code, url, created_at FROM short_urls
	WHERE code=$1`
	if err := r.db.QueryRow(ctx, sql, code).Scan(&res.ID, &res.Code, &res.URL, &res.CreatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return res, utils.NewAppErr(err.Error(), utils.ERR_OBJ_NOT_FOUND)
		}
		return res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
	}
	return res, nil
}

func (r urlRepository) List(ctx context.Context, req domain.ListShortURLReq) (int, []domain.ShortURL, utils.AppErr) {
	res := []domain.ShortURL{}
	count := 0

	offset := (req.Page - 1) * req.Limit

	csql := `
	SELECT COUNT(*) from short_urls`
	if err := r.db.QueryRow(ctx, csql).Scan(&count); err != nil {
		return 0, res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
	}

	sql := `
	SELECT id, code, url, created_at from short_urls
	LIMIT $1
	OFFSET $2`
	rows, err := r.db.Query(ctx, sql, req.Limit, offset)
	if err != nil {
		return 0, res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
	}
	if rows.Err() != nil {
		return 0, res, utils.NewAppErr(rows.Err().Error(), utils.ERR_UNKNOWN)
	}
	for rows.Next() {
		var s domain.ShortURL
		err := rows.Scan(&s.ID, &s.Code, &s.URL, &s.CreatedAt)
		if err != nil {
			return 0, res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
		}
		res = append(res, s)
	}

	return count, res, nil
}
