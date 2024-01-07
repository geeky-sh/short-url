package handlers

import (
	"net/http"
	"shorturl/domain"
	"shorturl/utils"
	"shorturl/utils/session"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type urlHandler struct {
	usc  domain.URLUsecase
	vl   *validator.Validate
	sess *session.Store
}

func NewURLhandler(s domain.URLUsecase, v *validator.Validate, sess *session.Store) urlHandler {
	return urlHandler{usc: s, vl: v, sess: sess}
}

func (h *urlHandler) Routes() http.Handler {
	r := chi.NewRouter()
	r.Post("/", h.Create)
	r.Get("/", h.List)
	r.Get("/{code}", h.Get)
	return r
}

// swagger:route POST / Url GenerateCode
// This API is used to create short url for the given url
func (h *urlHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := domain.URLCreateReq{}

	sessKey, err := r.Cookie(session.SESSION_KEY)
	if err != nil {
		utils.WriteMsgRes(w, http.StatusUnauthorized, err.Error())
		return
	}

	userID, err := h.sess.GetID(sessKey.Value)
	if err != nil {
		utils.WriteMsgRes(w, http.StatusUnauthorized, err.Error())
		return
	}

	if err := utils.JSNDecode(r, &req); err != nil {
		utils.WriteMsgRes(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.vl.Struct(req); err != nil {
		utils.WriteMsgRes(w, http.StatusBadRequest, err.Error())
		return
	}

	res, aerr := h.usc.Create(ctx, req.URL, userID)
	if aerr != nil {
		utils.WriteAppErrRes(w, aerr)
		return
	}

	utils.WriteRes(w, http.StatusCreated, res)
}

func (h *urlHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	code := chi.URLParam(r, "code")

	res, err := h.usc.Get(ctx, code)
	if err != nil {
		utils.WriteAppErrRes(w, err)
		return
	}

	http.Redirect(w, r, res, http.StatusPermanentRedirect)
}

func (h *urlHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if _, err := r.Cookie(session.SESSION_KEY); err != nil {
		utils.WriteMsgRes(w, http.StatusUnauthorized, err.Error())
		return
	}

	res, err := h.usc.List(ctx, domain.URLListReq{Page: 1, Limit: 20})
	if err != nil {
		utils.WriteAppErrRes(w, err)
		return
	}

	utils.WriteRes(w, http.StatusOK, res)
}
