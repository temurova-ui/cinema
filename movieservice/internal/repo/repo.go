package repository

import (
	"context"
	"errors"
	"github.com/temurova-ui/cinema/movieservice/internal/models"
	"github.com/temurova-ui/cinema/movieservice/movie"
	errs "github.com/temurova-ui/cinema/movieservice/pkg/errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) Repository {
	return Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, movie models.Movie) (int64, error) {
	const query = `
		INSERT INTO movies (title, description, duration, age_limit)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	var id int64
	err := r.db.QueryRow(
		ctx,
		query,
		movie.Title,
		movie.Description,
		movie.Duration,
		movie.AgeLimit,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) GetByID(ctx context.Context, id int64) (*movie.GetMovieResponse, error) {
	var myMovie movie.GetMovieResponse

	const query = `
		SELECT id, title, description, duration, age_limit, created_at::text 
		FROM movies
		WHERE id = $1`

	err := r.db.QueryRow(ctx, query, id).Scan(
		&myMovie.Id,
		&myMovie.Title,
		&myMovie.Description,
		&myMovie.Duration,
		&myMovie.AgeLimit,
		&myMovie.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.ErrMovieNotFound
		}

		return nil, err
	}

	return &myMovie, nil
}

func (r *Repository) List(ctx context.Context) (*movie.ListMovieResponse, error) {
	const query = `
		SELECT id, title, description, duration, age_limit, created_at::text
		FROM movies
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var myMovies movie.ListMovieResponse

	for rows.Next() {
		var myMovie movie.GetMovieResponse

		err = rows.Scan(
			&myMovie.Id,
			&myMovie.Title,
			&myMovie.Description,
			&myMovie.Duration,
			&myMovie.AgeLimit,
			&myMovie.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		myMovies.Movies = append(myMovies.Movies, &myMovie)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &myMovies, nil
}

func (r *Repository) Update(ctx context.Context, req *models.Movie) error {
	const query = `UPDATE movies 
					SET title = $2, description = $3, duration = $4, age_limit = $5 
					WHERE id = $1`

	result, err := r.db.Exec(ctx, query, req.ID, req.Title, req.Description, req.Duration, req.AgeLimit)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errs.ErrMovieNotFound
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, id int64) error {
	const query = `DELETE FROM movies WHERE id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errs.ErrMovieNotFound
	}

	return nil
}
