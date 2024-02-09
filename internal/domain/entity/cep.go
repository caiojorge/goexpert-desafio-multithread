package entity

import "github.com/google/uuid"

type Cep struct {
	ID           uuid.UUID `json:"id"`
	Cep          string    `json:"cep"`
	State        string    `json:"state"`
	City         string    `json:"city"`
	Neighborhood string    `json:"neighborhood"`
	Street       string    `json:"street"`
	Service      string    `json:"service"`
	ApiName      string    `json:"api_name"`
}
