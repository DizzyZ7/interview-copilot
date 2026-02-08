.PHONY: up down backend frontend cli test lint

up:
	docker compose up --build

down:
	docker compose down -v

backend:
	cd backend && go run cmd/server/main.go

frontend:
	cd frontend && npm run dev

cli:
	cd cli && go run cmd/copilot/main.go

test:
	cd backend && go test ./...

lint:
	echo "Add golangci-lint here"
