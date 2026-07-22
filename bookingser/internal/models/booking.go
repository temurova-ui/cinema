package models

import(
	"fmt"
)

type BookingStatus int32

const (
	StatusUnspecified  BookingStatus = 0
	StatusPending      BookingStatus = 1
	StatusConfirmed	   BookingStatus = 2
	StatusConcelled    BookingStatus = 3
)

type Booking struct {
	ID 		int
	User_ID int
	Movie_ID int
	Status   BookingStatus
}

func (b *Booking) Validate() error {
	if b.Movie_ID == 0 && b.User_ID == 0{
		return fmt.Errorf("Validation error: movei_id and user_id are required")
	}
	return nil
}