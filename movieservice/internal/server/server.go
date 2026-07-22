package server

import (
	"context"
	"errors"
	"github.com/temurova-ui/cinema/movieservice/internal/models"
	"github.com/temurova-ui/cinema/movieservice/internal/service"
	"github.com/temurova-ui/cinema/movieservice/movie"
	errs "github.com/temurova-ui/cinema/movieservice/pkg/errors"
	"github.com/temurova-ui/cinema/movieservice/pkg/logger"
)

type Server struct {
	movie.UnimplementedMovieServiceServer
	lg      logger.Logger
	service service.Service
}

func New(lg logger.Logger, service service.Service) *Server {
	return &Server{
		lg:      lg,
		service: service,
	}
}

func (s *Server) Create(ctx context.Context, req *movie.CreateMovieRequest) (*movie.CreateMovieResponse, error) {

	request := models.Movie{
		Title:       req.Title,
		Description: req.Description,
		Duration:    req.Duration,
		AgeLimit:    req.AgeLimit,
	}

	movieID, err := s.service.Create(ctx, request)
	if err != nil {
		if errors.Is(err, errs.ErrValidate) {
			return &movie.CreateMovieResponse{}, errs.ErrValidate
		}
		s.lg.Error("error from s.service.Create")
		return nil, err
	}

	return &movie.CreateMovieResponse{Id: movieID}, nil
}

func (s *Server) GetByID(ctx context.Context, req *movie.GetMovieRequest) (*movie.GetMovieResponse, error) {

	myMovie, err := s.service.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, errs.ErrMovieNotFound) {
			return nil, errs.ErrMovieNotFound
		}
		s.lg.Error("error from s.service.GetByID")
		return nil, err
	}

	return &movie.GetMovieResponse{
		Id:          myMovie.Id,
		Title:       myMovie.Title,
		Description: myMovie.Description,
		Duration:    myMovie.Duration,
		AgeLimit:    myMovie.AgeLimit,
		CreatedAt:   myMovie.CreatedAt,
	}, nil
}

func (s *Server) List(ctx context.Context, req *movie.ListMovieRequest) (*movie.ListMovieResponse, error) {
	myMovies, err := s.service.List(ctx)
	if err != nil {
		s.lg.Error("error from s.service.List")
		return nil, err
	}

	return &movie.ListMovieResponse{
		Movies: myMovies.Movies,
	}, nil
}

func (s *Server) Update(ctx context.Context, req *movie.UpdateMovieRequest) (*movie.UpdateMovieResponse, error) {

	response, err := s.service.Update(ctx, &models.Movie{
		ID:          req.Id,
		Title:       req.Title,
		Description: req.Description,
		AgeLimit:    req.AgeLimit,
		Duration:    req.Duration,
	})
	if err != nil {
		if errors.Is(err, errs.ErrMovieNotFound) {
			return nil, errs.ErrMovieNotFound
		}
		s.lg.Error("error from s.service.Update")
		return nil, err
	}

	return response, nil
}

func (s *Server) Delete(ctx context.Context, req *movie.DeleteMovieRequest) (*movie.DeleteMovieResponse, error) {

	movieID := req.Id

	response, err := s.service.Delete(ctx, movieID)
	if err != nil {
		if errors.Is(err, errs.ErrMovieNotFound) {
			return nil, errs.ErrMovieNotFound
		}
		s.lg.Error("error from s.service.Delete")
		return nil, err
	}

	return response, nil
}
