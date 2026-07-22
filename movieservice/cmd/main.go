package main

import (
	"context"
	"log"
	"github.com/temurova-ui/cinema/movieservice/internal/config"
	"github.com/temurova-ui/cinema/movieservice/internal/repo"
	"github.com/temurova-ui/cinema/movieservice/internal/server"
	"github.com/temurova-ui/cinema/movieservice/internal/service"
	"github.com/temurova-ui/cinema/movieservice/movie"
	"github.com/temurova-ui/cinema/movieservice/pkg/db"
	"github.com/temurova-ui/cinema/movieservice/pkg/logger"
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
		Port:     cfg.Port,
		User:     cfg.DBUser,
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
		log.Fatal("failed to listen: %w", err)
	}

	grpcServer := grpc.NewServer()

	movieRepo := repository.New(conn)

	movieService := service.New(movieRepo)

	movieServer := server.New(*lg, movieService)

	movie.RegisterMovieServiceServer(grpcServer, movieServer)
	
	reflection.Register(grpcServer)

	go func() {
		lg.Info("server listening", zap.String("addr", lis.Addr().String()))
		if err = grpcServer.Serve(lis); err != nil {
			log.Fatal("failed to serve: %w", err)
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
