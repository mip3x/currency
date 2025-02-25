package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"currency_app/internal/config"
	httpdelivery "currency_app/internal/delivery/http"
	// set alias to be sure that function is called not from net/http

	"currency_app/internal/adapter"
	"currency_app/internal/service"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	provider := adapter.NewCBRProvider(cfg)
	currencyService := service.NewCurrencyService(provider)
	router := httpdelivery.NewRouter(cfg, currencyService)

	port := cfg.Port
	if envPort, ok := os.LookupEnv("PORT"); ok {
		port = envPort
	}

	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(addr, router))
}
