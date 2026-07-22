package main

import (
	bookingv1 "github.com/temurova-ui/cinema/bookingser/bookingpb"
	"github.com/temurova-ui/cinema/bookingser/internal/config"
	"github.com/temurova-ui/cinema/bookingser/internal/repository"
	"github.com/temurova-ui/cinema/bookingser/internal/server"
	"github.com/temurova-ui/cinema/bookingser/internal/service"
	"github.com/temurova-ui/cinema/bookingser/pkg/db"
	"github.com/temurova-ui/cinema/bookingser/pkg/logger"
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	cfg, err := config.New("./config/config.env")
	if err != nil { 
		log.Fatal("config.New", err)
	}

	conn, err := db.New(db.Option{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		Booking:  cfg.DBBooking,
		Password: cfg.DBPassword,
		DBName:   cfg.DBName,
	})
	if err != nil {
		log.Fatal("failed to connect to db:", err)
	}
	defer conn.Close()

	lg, err := logger.New(true)
	if err != nil {
		log.Fatal("failed to create logger", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lis, err := net.Listen(cfg.NETWORK, cfg.ADDRESS)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	bookingRepo := repository.New(conn)

	bookingService := service.New(*bookingRepo, *lg)

	bookingServer := server.New(*lg, bookingService)

	bookingv1.RegisterBookingServiceServer(grpcServer, bookingServer)

	reflection.Register(grpcServer)

	go func() {
		lg.Info("server listening", zap.String("addr", lis.Addr().String()))
		if err = grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	lg.Info("shutting down server...")
	grpcServer.GracefulStop()
	lg.Info("server stopped")
}
