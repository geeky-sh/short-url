package usecases

import (
	"context"
	"shorturl/domain"
	"shorturl/utils"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(r domain.UserRepository) domain.UserUsecase {
	return userUsecase{repo: r}
}

func (r userUsecase) Create(ctx context.Context, req domain.UserCreateReq) (domain.UserResp, utils.AppErr) {
	res := domain.UserResp{}

	_, aerr := r.repo.GetByUsername(ctx, req.Username)
	if aerr == nil {
		return res, utils.NewAppErr("User name already exists", utils.ERR_ALREADY_EXISTS)
	}
	if aerr.ErrCode() != utils.ERR_OBJ_NOT_FOUND {
		return res, aerr
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
	}

	dbReq := req.DB(string(encryptedPassword))
	res, aerr = r.repo.Create(ctx, dbReq)
	if aerr != nil {
		return res, nil
	}
	return res, nil
}

func (r userUsecase) Login(ctx context.Context, req domain.UserLoginReq) (domain.UserResp, utils.AppErr) {
	res := domain.UserResp{}
	dbRes, err := r.repo.GetByUsername(ctx, req.Username)
	if err != nil {
		if err.ErrCode() == utils.ERR_OBJ_NOT_FOUND {
			return res, utils.NewAppErr("Invalid Login Credentails", utils.ERR_INVALID_CREDS)
		}
		return res, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbRes.EncryptedPassword), []byte(req.Password)); err != nil {
		return res, utils.NewAppErr("Invalid Login Credentails", utils.ERR_INVALID_CREDS)
	}

	return dbRes, nil
}
