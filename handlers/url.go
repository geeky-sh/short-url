package handlers

import (
	"net/http"
	"shorturl/entities"

	"github.com/go-chi/chi/v5"
)

type urlHandler struct {
	usc entities.URLUsecase
}

func NewURLhandler(s entities.URLUsecase) urlHandler {
	return urlHandler{usc: s}
}

func (h *urlHandler) Routes() http.Handler {
	r := chi.NewRouter()
	r.Post("/", h.Create)
	r.Get("/{id}", h.Get)
	return r
}

func (h *urlHandler) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *urlHandler) Get(w http.ResponseWriter, r *http.Request) {

}
