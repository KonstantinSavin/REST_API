FROM golang:latest AS builder

WORKDIR /music-lib

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o ./bin/music-lib ./cmd/app/main.go

CMD [ "./bin/music-lib" ]
