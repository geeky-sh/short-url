package handlers

import (
	"net/http"
	"shorturl/utils"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type healthHandler struct {
	db *pgxpool.Pool
}

func NewHealthHandler(db *pgxpool.Pool) healthHandler {
	return healthHandler{db}
}

func (h *healthHandler) Routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/health", h.HealthCheck)
	return r
}

func (h *healthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := h.db.Ping(ctx); err != nil {
		utils.WriteMsgRes(w, http.StatusInternalServerError, "Unable to connect to DB")
	}

	utils.WriteMsgRes(w, http.StatusOK, "success")
}
