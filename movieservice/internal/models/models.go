package models

import (
	"movieSer/pkg/errors"
	"time"
)

type Movie struct {
	ID          int64
	Title       string
	Description string
	Duration    int32
	AgeLimit    int32
	CreatedAt   time.Time
}

func (m *Movie) Validate() error {
	if len(m.Title) == 0 || len(m.Description) == 0 || m.AgeLimit < 1 || m.Duration < 1 {
		return errors.ErrValidate
	}
	return nil
}
