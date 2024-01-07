package repositories

import (
	"context"
	"errors"
	"shorturl/domain"
	"shorturl/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) domain.UserRepository {
	return userRepository{db: db}
}

func (r userRepository) Create(ctx context.Context, req domain.UserDB) (domain.UserDB, utils.AppErr) {
	res := domain.UserDB{}

	sql := `
	INSERT INTO users (username, encrypted_password, created_at)
	VALUES ($1, $2, $3) RETURNING id
	`
	_, err := r.db.Exec(ctx, sql, req.Username, req.EncryptedPassword, req.CreatedAt)
	if err != nil {
		return res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
	}
	return req, nil
}

func (r userRepository) GetByUsername(ctx context.Context, username string) (domain.UserDB, utils.AppErr) {
	res := domain.UserDB{}

	sql := `
	SELECT id, username, encrypted_password FROM users
	WHERE username=$1
	`
	if err := r.db.QueryRow(ctx, sql, username).Scan(&res.ID, &res.Username, &res.EncryptedPassword); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return res, utils.NewAppErr(err.Error(), utils.ERR_OBJ_NOT_FOUND)
		}
		return res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
	}
	return res, nil
}
