gen:
	go run ./cmd/gen/. -conf .env.development
dev:
	go run main.go
prod:
	go run main.go -conf .env.production