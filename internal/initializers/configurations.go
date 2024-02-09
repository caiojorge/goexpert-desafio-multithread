package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

type Configurations struct{}

func (c *Configurations) Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
