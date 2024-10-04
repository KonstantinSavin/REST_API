build:
		go build -v ./cmd/app

test:
		go test -v -timeout 30s ./...

up:
	docker-compose up --force-recreate

down:
	docker-compose down
	docker image rm music-lib-app

.DEFAULT_GOAL := build