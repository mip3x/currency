package http

import (
	"net/http"

	"currency_app/internal/config"
	"currency_app/internal/service"
)

func NewRouter(cfg *config.Config, currencyService *service.CurrencyService) http.Handler {
	h := NewHandler(cfg, currencyService)
	mux := http.NewServeMux()

	mux.HandleFunc("/info", h.InfoHandler)
	mux.HandleFunc("/info/currency", h.CurrencyHandler)

	return mux
}
