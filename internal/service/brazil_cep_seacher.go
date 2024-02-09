package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/caiojorge/goexpert-desafio-multithread/internal/domain/dto"
	"github.com/caiojorge/goexpert-desafio-multithread/internal/domain/entity"
	"github.com/caiojorge/goexpert-desafio-multithread/internal/infra"
)

type BrazilCepSearcher struct{}

func (b *BrazilCepSearcher) GetCep(ctx context.Context, cep string) (interface{}, error) {
	var url string
	urlEnvBrazilCep := os.Getenv("BRASILCEP")
	if urlEnvBrazilCep != "" {
		url = fmt.Sprintf("%s/%s", urlEnvBrazilCep, cep)
	} else {
		url = fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	var cepResponse dto.BrazilCep
	if err := json.NewDecoder(resp.Body).Decode(&cepResponse); err != nil {
		return nil, err
	}

	infra := infra.UuidGenerator{}
	id := infra.Generate()

	result := entity.Cep{
		ID:           id,
		Cep:          cepResponse.Cep,
		State:        cepResponse.State,
		City:         cepResponse.City,
		Neighborhood: cepResponse.Neighborhood,
		Street:       cepResponse.Street,
		Service:      cepResponse.Service,
		ApiName:      "BrasilAPI",
	}

	return result, nil
}
