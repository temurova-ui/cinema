FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /bookingservice ./bookingser/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /bookingservice /app/bookingservice
COPY --from=builder /app/bookingser/config /app/config

EXPOSE 50053

CMD ["/app/bookingservice"]
