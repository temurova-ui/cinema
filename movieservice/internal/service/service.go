package service

import (
	"context"
	"fmt"
	"movieSer/internal/models"
	"movieSer/internal/repo"
	"movieSer/movie"
	"movieSer/pkg/errors"
)

type Service struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, request models.Movie) (int64, error) {

	err := request.Validate()
	if err != nil {
		return 0, errors.ErrValidate
	}

	id, err := s.repo.Create(ctx, request)
	if err != nil {
		return 0, fmt.Errorf("error from s.repo.Create: %w", err)
	}

	return id, nil
}

func (s *Service) GetByID(ctx context.Context, id int64) (*movie.GetMovieResponse, error) {

	myMovie, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return myMovie, nil
}

func (s *Service) List(ctx context.Context) (*movie.ListMovieResponse, error) {

	myMovies, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	return myMovies, nil
}

func (s *Service) Update(ctx context.Context, req *models.Movie) (*movie.UpdateMovieResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, errors.ErrValidate
	}

	err = s.repo.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	return &movie.UpdateMovieResponse{
		Message: "Movie successfully updated",
		Code:    0,
	}, nil
}

func (s *Service) Delete(ctx context.Context, id int64) (*movie.DeleteMovieResponse, error) {

	if id < 1 {
		return nil, errors.ErrValidate
	}

	err := s.repo.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return &movie.DeleteMovieResponse{
		Message: "Movie successfully deleted",
		Code:    0,
	}, nil
}
