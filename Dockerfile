FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o analytics_service ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/analytics_service .

COPY internal/config /root/internal/config
COPY .env .

CMD ["./analytics_service"]
