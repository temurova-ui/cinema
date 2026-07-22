package db

import(
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Option struct {
	Host	 string
	Port     int
	Booking    string
	Password string
	DBName   string
}

func New(o Option) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("host=%s port=%d movie=%s password=%s dbname=%s",
		o.Host, o.Port, o.Booking, o.Password, o.DBName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	if err = pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}
	return pool, nil
}