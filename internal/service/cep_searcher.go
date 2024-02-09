package service

import "context"

type CepSearcher interface {
	GetCep(ctx context.Context, cep string) (interface{}, error)
}

type CepService struct {
	Searcher CepSearcher
	Ctx      context.Context
	Cep      string
}

func (c *CepService) GetCep() (interface{}, error) {
	return c.Searcher.GetCep(c.Ctx, c.Cep)
}
