package dbmodel

import (
	"gorm.io/gorm"
	"time"
)

type Property struct {
	ID                 int
	Name               string
	Zip                int
	Address            string
	CurrentPricePerDay float64
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt
}

type Booking struct {
	ID          int
	UserID      int
	PropertyID  int
	StartDate   time.Time
	EndDate     time.Time
	PricePerDay float64
	Paid        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
