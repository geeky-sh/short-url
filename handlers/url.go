package handlers

import (
	"fmt"
	"net/http"
	"shorturl/domain"
	"shorturl/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type urlHandler struct {
	usc domain.URLUsecase
	vl  *validator.Validate
}

func NewURLhandler(s domain.URLUsecase, v *validator.Validate) urlHandler {
	return urlHandler{usc: s, vl: v}
}

func (h *urlHandler) Routes() http.Handler {
	r := chi.NewRouter()
	r.Post("/", h.Create)
	r.Get("/{code}", h.Get)
	return r
}

func (h *urlHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := domain.CreateShortURL{}

	if err := utils.JSNDecode(r, &req); err != nil {
		utils.WriteMsgRes(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.vl.Struct(req); err != nil {
		utils.WriteMsgRes(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.usc.Create(ctx, req.URL)
	if err != nil {
		utils.WriteAppErrRes(w, err)
		return
	}

	utils.WriteMsgRes(w, http.StatusCreated, res)
}

func (h *urlHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	code := chi.URLParam(r, "code")

	res, err := h.usc.Get(ctx, code)
	if err != nil {
		utils.WriteAppErrRes(w, err)
		return
	}

	fmt.Println(res)

	http.Redirect(w, r, res, http.StatusPermanentRedirect)
}
