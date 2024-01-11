package handlers

import (
	"database/sql"
	"net/http"
	"shorturl/utils"

	"github.com/go-chi/chi/v5"
)

type healthHandler struct {
	db *sql.DB
}

func NewHealthHandler(db *sql.DB) healthHandler {
	return healthHandler{db}
}

func (h *healthHandler) Routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/health", h.HealthCheck)
	return r
}

func (h *healthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	if err := h.db.Ping(); err != nil {
		utils.WriteMsgRes(w, http.StatusInternalServerError, "Unable to connect to DB")
	}

	utils.WriteMsgRes(w, http.StatusOK, "success")
}
