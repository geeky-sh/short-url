package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"shorturl/domain"
	"shorturl/utils"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return userRepository{db: db}
}

func (r userRepository) Create(ctx context.Context, req domain.UserDB) (domain.UserDB, utils.AppErr) {
	res := domain.UserDB{}

	sql := `
	INSERT INTO users (username, encrypted_password, created_at)
	VALUES ($1, $2, $3)
	`
	_, err := r.db.ExecContext(ctx, sql, req.Username, req.EncryptedPassword, req.CreatedAt)
	if err != nil {
		fmt.Println(err.Error())
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
	if err := r.db.QueryRowContext(ctx, sql, username).Scan(&res.ID, &res.Username, &res.EncryptedPassword); err != nil {
		fmt.Println(err.Error())
		if err.Error() == utils.ERR_MSG_NO_OBJECT_FOUND {
			return res, utils.NewAppErr(err.Error(), utils.ERR_OBJ_NOT_FOUND)
		}
		return res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
	}
	return res, nil
}
