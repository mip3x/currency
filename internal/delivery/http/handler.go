package http

import (
	"currency_app/internal/config"
	"currency_app/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Handler struct {
	cfg             *config.Config
	currencyService *service.CurrencyService
}

func NewHandler(cfg *config.Config, cs *service.CurrencyService) *Handler {
	return &Handler{
		cfg:             cfg,
		currencyService: cs,
	}
}

func (h *Handler) InfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"version": "%s", "service": "%s", "author": "%s"}`, h.cfg.Version, h.cfg.Service, h.cfg.Author)
}

func (h *Handler) CurrencyHandler(w http.ResponseWriter, r *http.Request) {
	dateParam := r.URL.Query().Get("date")
	if dateParam != "" {
		parsedDate, err := time.Parse("2006-01-02", dateParam)
		if err != nil {
			http.Error(w, "Invalid date format: use YYYY-MM-DD", http.StatusBadRequest)
			return
		}
		dateParam = parsedDate.Format("02/01/2006")
	}

	currencyQuery := r.URL.Query().Get("currency")
	if currencyQuery != "" {
		currencyQuery = strings.ToUpper(currencyQuery)
	}

	data, err := h.currencyService.GetCurrencyData(dateParam, currencyQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]any{
		"service": h.cfg.Service,
		"data":    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResp, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Write(jsonResp)
}
