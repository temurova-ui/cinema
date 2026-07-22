FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /movieservice ./movieservice/cmd/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /movieservice /app/movieservice
COPY --from=builder /app/movieservice/config /app/config

EXPOSE 50052

CMD ["/app/movieservice"]
