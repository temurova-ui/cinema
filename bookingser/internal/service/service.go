package service

import (
	"github.com/temurova-ui/cinema/bookingser/internal/models"
	"github.com/temurova-ui/cinema/bookingser/internal/repository"
	"github.com/temurova-ui/cinema/bookingser/pkg/logger"
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"
)

type Service struct {
	repo repository.Repository
	log *logger.Logger
}

func New(repo repository.Repository, log logger.Logger) Service {
	return Service{repo: repo,
	log: &log,
	}
}

func (s *Service) Create(ctx context.Context, request models.Booking) (int, error) {
	err := request.Validate()
	if err != nil {
		return 0, fmt.Errorf("Validation error")
	}

	BookingID, err := s.repo.Create(ctx, models.Booking{
		User_ID: request.User_ID,
		Movie_ID: request.Movie_ID,
		Status: request.Status,
	})
	if err != nil {
		return 0, fmt.Errorf("error from s.repo.Add")
	}

	return BookingID, nil
}

func (s *Service) GetBooking(ctx context.Context, id int) (models.Booking, error) {
	booking, err := s.repo.Get(ctx, id)
	if err != nil {
		return models.Booking{}, err
	}
	return booking, nil
}

func (s *Service) GetUserBooking(ctx context.Context, UserId int) ([]*models.Booking, error) {
	booking, err := s.repo.GetByUserID(ctx, UserId)
	if err != nil {
		s.log.Error("massage failed", zap.Int("UserId", UserId), zap.Error(err))
		return nil, fmt.Errorf("%w: %v", errors.New("Error"), err)
	}
	return booking, nil
}

func (s *Service) CancelBooking(id int) error {
	if id == 0 {
		return errors.New("invalid id")
	}
	return s.repo.CancelBooking(context.Background(), int(id))
}
