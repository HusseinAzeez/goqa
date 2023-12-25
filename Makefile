build:
	@go build -o bin/gofuneral

run: build
	@./bin/gofuneral

test:
	@go test -v ./...

format:
	@go fmt ./...

