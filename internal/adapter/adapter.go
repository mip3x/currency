package adapter

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"currency_app/internal/config"
	"currency_app/internal/domain"

	"golang.org/x/net/html/charset"
)

type CBRProvider struct {
	config *config.Config
	client *http.Client
}

func NewCBRProvider(cfg *config.Config) *CBRProvider {
	return &CBRProvider{
		config: cfg,
		client: &http.Client{},
	}
}

func (p *CBRProvider) GetCurrencies(date string) (map[string]float64, error) {
	url := p.config.URL + "?date_req=" + date
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("User-Agent", p.config.UserAgent)

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch currency data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("CBR returned error: %s", resp.Status)
	}

	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel

	var vc domain.ValCurs
	if err := decoder.Decode(&vc); err != nil {
		return nil, fmt.Errorf("failed to parse XML: %v", err)
	}

	data := make(map[string]float64)
	for _, v := range vc.Valutes {
		data[v.CharCode] = v.Value
	}

	return data, nil
}
