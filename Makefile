build:
		go build -v ./cmd/app

test:
		go test -v -timeout 30s ./...

up:
	docker-compose up --force-recreate

.DEFAULT_GOAL := build