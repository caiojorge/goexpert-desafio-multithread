install:
	go get github.com/gofiber/fiber/v2
	go get github.com/joho/godotenv
	go get github.com/google/uuid

run-server:
	go run cmd/server/main.go

run-client:
	go run cmd/client/main.go -cep=01311200

run-client-via:
	go run cmd/client/main.go -cep=01311200 -searcher=via

run-client-brazil:
	go run cmd/client/main.go -cep=01311200 -searcher=brazil