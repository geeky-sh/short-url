package repositories

import (
	"context"
	"database/sql"
	"errors"
	"shorturl/domain"
	"shorturl/utils"

	"github.com/jackc/pgx/v5"
)

const TABLE_NAME = "urls"

type urlRepository struct {
	db *sql.DB
	tb string
}

func NewURLRepository(db *sql.DB) domain.URLRepository {
	return urlRepository{db, TABLE_NAME}
}

func (r urlRepository) Create(ctx context.Context, req domain.URLDB) (domain.URLDB, utils.AppErr) {
	res := domain.URLDB{}

	sql := `
	INSERT INTO urls (code, url, created_at, user_id)
	VALUES ($1, $2, $3, $4) RETURNING id`
	_, err := r.db.ExecContext(ctx, sql, req.Code, req.URL, req.CreatedAt, req.UserID)
	if err != nil {
		return res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
	}

	return req, nil
}

func (r urlRepository) GetByURL(ctx context.Context, url string) (domain.URLDB, utils.AppErr) {
	res := domain.URLDB{}

	sql := `
	SELECT id, code, url, created_at FROM urls
	WHERE url=$1`
	if err := r.db.QueryRowContext(ctx, sql, url).Scan(res.ID, res.Code, res.URL, res.CreatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return res, utils.NewAppErr(err.Error(), utils.ERR_OBJ_NOT_FOUND)
		}
		return res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
	}
	return res, nil
}

func (r urlRepository) GetByCode(ctx context.Context, code string) (domain.URLDB, utils.AppErr) {
	res := domain.URLDB{}

	sql := `
	SELECT id, code, url, created_at FROM urls
	WHERE code=$1`
	if err := r.db.QueryRowContext(ctx, sql, code).Scan(&res.ID, &res.Code, &res.URL, &res.CreatedAt); err != nil {
		if err.Error() == utils.ERR_MSG_NO_OBJECT_FOUND {
			return res, utils.NewAppErr(err.Error(), utils.ERR_OBJ_NOT_FOUND)
		}
		return res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
	}
	return res, nil
}

func (r urlRepository) List(ctx context.Context, req domain.URLListReq) (int, []domain.URLDB, utils.AppErr) {
	res := []domain.URLDB{}
	count := 0

	offset := (req.Page - 1) * req.Limit

	csql := `
	SELECT COUNT(*) from urls`
	if err := r.db.QueryRowContext(ctx, csql).Scan(&count); err != nil {
		return 0, res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
	}

	sql := `
	SELECT id, user_id, code, url, created_at from urls
	LIMIT $1
	OFFSET $2`

	rows, err := r.db.QueryContext(ctx, sql, req.Limit, offset)
	if err != nil {
		return 0, res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
	}

	for rows.Next() {
		r := domain.URLDB{}
		if err := rows.Scan(&r.ID, &r.UserID, &r.Code, &r.URL, &r.CreatedAt); err != nil {
			return 0, res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
		}
		res = append(res, r)
	}

	return count, res, nil
}
