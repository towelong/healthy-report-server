gen:
	go run ./cmd/gen/.
dev:
	go run main.go
prod:
	go run main.go -e .env.production