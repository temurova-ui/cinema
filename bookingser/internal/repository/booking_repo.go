package repository

import (
	"github.com/temurova-ui/cinema/bookingser/internal/models"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, request models.Booking) (int, error) {
	const query = `
	INSERT INTO mv_booking (user_id, movie_id, status)
	VALUES ($1, $2, $3)
	RETURNING id`

	var id int
	err := r.db.QueryRow(ctx, query, request.User_ID, request.Movie_ID, request.Status).Scan(&id)
	if err != nil {
		return 0, errors.New("error from r.db.Exec")
	}

	return id, nil
}

func (r *Repository) Get(ctx context.Context, id int) (models.Booking, error) {
	const query = `SELECT id, user_id, movie_id, status FROM mv_booking WHERE id = $1`

	var book models.Booking
	err := r.db.QueryRow(ctx, query, id).Scan(
		&book.ID,
		&book.User_ID,
		&book.Movie_ID,
		&book.Status,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Booking{}, fmt.Errorf("Booking not found")
		}
		return models.Booking{}, err
	}
	return book, err
}

func (r *Repository) GetByUserID(ctx context.Context, UserID int) ([]*models.Booking, error) {
	const query = `SELECT id, user_id, movie_id, status FROM mv_booking WHERE user_id = $1`
	rows, err := r.db.Query(ctx, query, UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var booking []*models.Booking
	for rows.Next() {
		b := &models.Booking{}
		if err := rows.Scan(&b.ID, &b.User_ID, &b.Movie_ID, &b.Status); err != nil {
			return nil, err
		}
		booking = append(booking, b)
	}
	return booking, nil
}

func (r *Repository) CancelBooking(ctx context.Context, id int) error {
	const query = `UPDATE mv_booking SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}