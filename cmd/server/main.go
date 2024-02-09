package main

import (
	"github.com/caiojorge/goexpert-desafio-multithread/internal/handlers"
	"github.com/caiojorge/goexpert-desafio-multithread/internal/initializers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config := initializers.Configurations{}
	config.Load()

	h := handlers.Handlers{}

	app := fiber.New()

	// criar contexto com timeout
	// abrir 2 canais
	// buscar o cep via brazil_cep com o canal 1
	// buscar o cep via viacep com o canal 2
	// fazer um select p esperar pelo canal que for preenchido primero
	// retornar o cep que foi preenchido primeiro

	app.Get("/cep", h.GetCep)

	app.Listen(":3000")
}
