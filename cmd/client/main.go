package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/caiojorge/goexpert-desafio-multithread/internal/domain/entity"
	"github.com/caiojorge/goexpert-desafio-multithread/internal/initializers"
	"github.com/caiojorge/goexpert-desafio-multithread/internal/service"
)

func init() {
	c := initializers.Configurations{}
	c.Load()
}

func main() {
	cep := flag.String("cep", "88034050", "Digite um CEP válido")
	searcher := flag.String("searcher", "", "Se quiser rodar um searcher específico, digite o nome dele")

	flag.Parse()

	log.Default().Println("CEP:", *cep)
	log.Default().Println("Searcher:", *searcher)

	c1 := make(chan entity.Cep)
	c2 := make(chan entity.Cep)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if *searcher != "" {
		switch *searcher {
		case "brazil":
			log.Default().Println("BrazilCep")
			brazilCepInfo, err := GetBrazilCEPInfo(ctx, *cep)
			if err != nil {
				fmt.Println("Error fetching CEP info:", err)
				return
			}
			fmt.Printf("CEP Info: %+v\n", brazilCepInfo)
			return
		case "via":
			log.Default().Println("ViaCep")
			viaCepInfo, err := GetViaCEPInfo(ctx, *cep)
			if err != nil {
				fmt.Println("Error fetching CEP info:", err)
				return
			}

			fmt.Printf("CEP Info: %+v\n", viaCepInfo)
			return
		}
	}

	go func() {
		brazilCepInfo, err := GetBrazilCEPInfo(ctx, *cep)
		if err != nil {
			fmt.Println("Error fetching CEP info:", err)
			return
		}

		//fmt.Printf("CEP Info: %+v\n", brazilCepInfo)
		c1 <- *brazilCepInfo
	}()

	go func() {
		viaCepInfo, err := GetViaCEPInfo(ctx, *cep)
		if err != nil {
			fmt.Println("Error fetching CEP info:", err)
			return
		}

		//fmt.Printf("CEP Info: %+v\n", viaCepInfo)
		c2 <- *viaCepInfo
	}()

Loop:
	for {
		select {
		case msg := <-c1:
			fmt.Printf("Received from BrazilCep: ID: %s - %s - %s\n", msg.ApiName, msg.ID.String(), msg.Street)
			break Loop
		case msg := <-c2:
			fmt.Printf("Received from ViaCep: ID: %s - %s - %s\n", msg.ApiName, msg.ID.String(), msg.Street)
			break Loop
		case <-time.After(time.Second * 2):
			println("timeout")
			break Loop
		}
	}
}

func GetBrazilCEPInfo(ctx context.Context, cep string) (*entity.Cep, error) {

	s := service.CepService{
		Searcher: &service.BrazilCepSearcher{},
		Ctx:      ctx,
		Cep:      cep,
	}

	result, err := s.GetCep()
	if err != nil {
		return nil, err
	}

	cepEntity := result.(entity.Cep)

	return &cepEntity, nil
}

func GetViaCEPInfo(ctx context.Context, cep string) (*entity.Cep, error) {

	s := service.CepService{
		Searcher: &service.ViaCepSearcher{},
		Ctx:      ctx,
		Cep:      cep,
	}

	result, err := s.GetCep()
	if err != nil {
		return nil, err
	}

	cepEntity := result.(entity.Cep)

	return &cepEntity, nil
}
