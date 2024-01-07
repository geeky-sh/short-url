package domain

import (
	"context"
	"shorturl/utils"
	"time"
)

type UserDB struct {
	ID                uint       `json:"-"`
	Username          string     `json:"username"`
	EncryptedPassword string     `json:"-"`
	CreatedAt         time.Time  `json:"created_at"`
	LastLoggedInAt    *time.Time `json:"last_logged_in_at"`
}

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserCreateReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResp = UserDB

type UserUsecase interface {
	Create(ctx context.Context, req UserCreateReq) (UserResp, utils.AppErr)
	Login(ctx context.Context, req UserLoginReq) (UserResp, utils.AppErr)
}

type UserRepository interface {
	Create(ctx context.Context, req UserDB) (UserDB, utils.AppErr)
	GetByUsername(ctx context.Context, username string) (UserDB, utils.AppErr)
}

func (r UserCreateReq) DB(encryptedPass string) UserDB {
	return UserDB{Username: r.Username, EncryptedPassword: encryptedPass, CreatedAt: time.Now()}
}
