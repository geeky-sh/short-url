package handlers

import (
	"net/http"
	"shorturl/domain"
	"shorturl/utils"
	"shorturl/utils/session"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	usc  domain.UserUsecase
	vl   *validator.Validate
	sess *session.Store
}

func NewUserHandler(s domain.UserUsecase, v *validator.Validate, sess *session.Store) userHandler {
	return userHandler{usc: s, vl: v, sess: sess}
}

func (h *userHandler) Routes() http.Handler {
	r := chi.NewRouter()
	r.Post("/", h.Create)
	r.Post("/login", h.Login)
	return r
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := domain.UserCreateReq{}

	if err := utils.JSNDecode(r, &req); err != nil {
		utils.WriteMsgRes(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.vl.Struct(req); err != nil {
		utils.WriteMsgRes(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.usc.Create(ctx, req)
	if err != nil {
		utils.WriteAppErrRes(w, err)
		return
	}

	utils.WriteRes(w, http.StatusCreated, res)
}

func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := domain.UserLoginReq{}

	if err := utils.JSNDecode(r, &req); err != nil {
		utils.WriteMsgRes(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.vl.Struct(req); err != nil {
		utils.WriteMsgRes(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.usc.Login(ctx, req)
	if err != nil {
		utils.WriteAppErrRes(w, err)
		return
	}

	sessKey := h.sess.Create(res.ID, res.Username)

	utils.WriteRes(w, http.StatusOK, map[string]string{"access_token": sessKey})
}
