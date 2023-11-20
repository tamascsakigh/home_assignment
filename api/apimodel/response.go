package apimodel

import "time"

type GetFreePropertiesResponse struct {
	ID          int
	Name        string
	Zip         int
	Address     string
	PricePerDay float64
}

type GetBookingHistoryResponse struct {
	ID          int
	Name        string
	Zip         int
	Address     string
	PricePerDay float64
	StartDate   time.Time
	EndDate     time.Time
	Days        int
	Price       float64
}
