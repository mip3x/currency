package service

import (
	"fmt"

	"currency_app/internal/domain"
)

type CurrencyService struct {
	provider domain.CurrencyProvider
}

func NewCurrencyService(provider domain.CurrencyProvider) *CurrencyService {
	return &CurrencyService{
		provider: provider,
	}
}

func (s *CurrencyService) GetCurrencyData(date, currencyQuery string) (map[string]float64, error) {
	data, err := s.provider.GetCurrencies(date)
	if err != nil {
		return nil, fmt.Errorf("error getting currencies: %v", err)
	}

	if currencyQuery != "" {
		if val, ok := data[currencyQuery]; ok {
			return map[string]float64{currencyQuery: val}, nil
		}
		return nil, fmt.Errorf("currency %s not found", currencyQuery)
	}
	return data, nil
}
