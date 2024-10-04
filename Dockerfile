FROM golang:1.22-alpine AS builder
WORKDIR /music-lib
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o ./bin/app ./cmd/app/main.go

FROM alpine
COPY --from=builder /music-lib/bin/app /
COPY --from=builder /music-lib/migrations /migrations
COPY .env /

CMD [ "/app" ]
