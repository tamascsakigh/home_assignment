package apimodel

import "time"

type GetFreePropertiesRequest struct {
	MinPrice  float64
	MaxPrice  float64
	StartDate time.Time `validate:"required"`
	EndDate   time.Time `validate:"required"`
}

type CreateBookingRequest struct {
	PropertyID int       `validate:"required"`
	StartDate  time.Time `validate:"required"`
	EndDate    time.Time `validate:"required"`
}

type DeleteBookingRequest struct {
	BookingID int `validate:"required"`
}

type GetBookingHistoryRequest struct {
	MinPrice  float64
	MaxPrice  float64
	StartDate time.Time
	EndDate   time.Time
	Print     bool
}

type PayRequest struct {
	BookingID int `validate:"required"`
}
