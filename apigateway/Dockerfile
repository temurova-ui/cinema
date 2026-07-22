FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /apigateway ./apigateway/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /apigateway /app/apigateway
COPY --from=builder /app/apigateway/config /app/config

EXPOSE 8080

CMD ["/app/apigateway"]
