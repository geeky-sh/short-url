package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

type healthHandler struct {
	db *pgx.Conn
}

func NewHealthHandler(db *pgx.Conn) healthHandler {
	return healthHandler{db}
}

func (h *healthHandler) Routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/health", h.HealthCheck)
	return r
}

func (h *healthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	en := json.NewEncoder(w)

	_, err := h.db.Exec(ctx, ";")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf("unable to connect to DB %v\n", err)
		en.Encode(map[string]string{"msg": msg})
	}

	w.WriteHeader(http.StatusOK)
	en.Encode(map[string]string{"success": "ok"})
}
