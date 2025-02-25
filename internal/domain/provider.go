package domain

type CurrencyProvider interface {
	GetCurrencies(date string) (map[string]float64, error)
}
