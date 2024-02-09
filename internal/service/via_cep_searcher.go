package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/caiojorge/goexpert-desafio-multithread/internal/domain/dto"
	"github.com/caiojorge/goexpert-desafio-multithread/internal/domain/entity"
	"github.com/caiojorge/goexpert-desafio-multithread/internal/infra"
)

type ViaCepSearcher struct{}

func (v *ViaCepSearcher) GetCep(ctx context.Context, cep string) (interface{}, error) {

	var url string
	urlEnvViaCep := os.Getenv("VIACEP")
	if urlEnvViaCep != "" {
		url = strings.Replace(urlEnvViaCep, "{cep}", cep, -1)
	} else {
		url = fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

	}

	//url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

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

	var cepResponse dto.ViaCep
	if err := json.NewDecoder(resp.Body).Decode(&cepResponse); err != nil {
		return nil, err
	}

	infra := infra.UuidGenerator{}
	id := infra.Generate()

	result := entity.Cep{
		ID:           id,
		Cep:          cepResponse.Cep,
		State:        cepResponse.Uf,
		City:         cepResponse.Localidade,
		Neighborhood: cepResponse.Bairro,
		Street:       cepResponse.Logradouro,
		Service:      "via_cep",
		ApiName:      "ViaCEP",
	}

	return result, nil
}
