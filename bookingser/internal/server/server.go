package server

import (
	bookingv1 "github.com/temurova-ui/cinema/bookingser/bookingpb"
	"github.com/temurova-ui/cinema/bookingser/internal/models"
	"github.com/temurova-ui/cinema/bookingser/internal/service"
	"github.com/temurova-ui/cinema/bookingser/pkg/logger"
	"context"
	"errors"
)

type Server struct {
	bookingv1.UnimplementedBookingServiceServer
	lg      logger.Logger
	service service.Service
}

func New(lg logger.Logger, service service.Service) *Server {
	return &Server{
		lg:      lg,
		service: service,
	}
}

func (s *Server) Create(ctx context.Context, req *bookingv1.CreateBookingRequest) (*bookingv1.CreateBookingResponse, error) {

    request := models.Booking{
        User_ID:  int(req.UserId),
        Movie_ID: int(req.MovieId),
    }

    bookingID, err := s.service.Create(ctx, request)
    if err != nil {
        return nil, err
    }

    return &bookingv1.CreateBookingResponse{
        Booking: &bookingv1.Booking{
            Id:      int32(bookingID),   
            UserId:  req.UserId,
            MovieId: req.MovieId,
        },
    }, nil
}

func (s *Server) GetBooking(ctx context.Context, req *bookingv1.GetBookingRequest) (*bookingv1.GetBookingResponse, error) {
	id := req.Id
	if id == 0 {	
		return &bookingv1.GetBookingResponse{}, errors.New("invalid id")
	}

	BookingFromDB, err := s.service.GetBooking(ctx, int(id))
	if err != nil {
		return nil, err
	}
	return &bookingv1.GetBookingResponse{
		  Booking: &bookingv1.Booking{
		Id:    req.Id,
		UserId: int32(BookingFromDB.User_ID),
		MovieId: int32(BookingFromDB.Movie_ID),
		Status: bookingv1.BookingStatus(BookingFromDB.Status),
		},
	}, nil
}
