run:
	go run cmd/server/main.go

test:
	go test ./...

build:
	go build -o server cmd/server/main.go

lint:
	golangci-lint run || echo "golangci-lint not installed"
