package handler

import (
	"errors"
	"github.com/tamascsakigh/home_assignment/api/apimodel"
	"github.com/tamascsakigh/home_assignment/database/dbmodel"
	"time"
)

var ErrNoPermission = errors.New("do not have permission")
var ErrInvalidDates = errors.New("invalid dates")
var ErrDateConflict = errors.New("accommodation already booked")
var ErrPaid = errors.New("booking already paid")

type Repository interface {
	GetFreeProperties(minPrice, maxPrice float64, startDate, endDate time.Time) ([]dbmodel.Property, error)
	GetPropertyBookings(ID int, startDate, endDate time.Time) ([]dbmodel.Booking, error)
	GetBookingHistory(userID int, minPrice, maxPrice float64, startDate, endDate time.Time) ([]dbmodel.Booking, error)
	GetPropertiesUnscoped(IDs []int) ([]dbmodel.Property, error)
	GetProperty(ID int) (dbmodel.Property, error)
	GetBooking(ID int) (dbmodel.Booking, error)
	CreateBooking(userID, propertyID int, startDate, endDate time.Time, pricePerDay float64) error
	DeleteBooking(ID int) error
}

type Handler struct {
	repository Repository
}

func NewHandler(repository Repository) *Handler { return &Handler{repository: repository} }

func (h *Handler) GetFreeProperties(request apimodel.GetFreePropertiesRequest) (response []apimodel.GetFreePropertiesResponse, err error) {

	// Check request dates
	if request.StartDate.Before(time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location())) || request.StartDate.After(request.EndDate) {
		return response, ErrInvalidDates
	}

	// Get free properties from database
	properties, err := h.repository.GetFreeProperties(request.MinPrice, request.MaxPrice, request.StartDate, request.EndDate)

	// Create response model array
	for _, p := range properties {
		response = append(response, apimodel.GetFreePropertiesResponse{
			ID:          p.ID,
			Name:        p.Name,
			Zip:         p.Zip,
			Address:     p.Address,
			PricePerDay: p.CurrentPricePerDay,
		})
	}
	return
}

func (h *Handler) CreateBooking(userID int, request apimodel.CreateBookingRequest) error {

	// Check request dates
	if request.StartDate.Before(time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location())) || request.StartDate.After(request.EndDate) {
		return ErrInvalidDates
	}

	// Check bookings for the property at that time
	if bookings, err := h.repository.GetPropertyBookings(request.PropertyID, request.StartDate, request.EndDate); err != nil {
		return err
	} else if len(bookings) > 0 {
		return ErrDateConflict
	}

	// Get property data from database
	property, err := h.repository.GetProperty(request.PropertyID)
	if err != nil {
		return err
	}

	// Save booking entity
	return h.repository.CreateBooking(userID, request.PropertyID, request.StartDate, request.EndDate, property.CurrentPricePerDay)
}

func (h *Handler) DeleteBooking(userID int, request apimodel.DeleteBookingRequest) error {

	// Check booking is existing and permission to delete
	if booking, err := h.repository.GetBooking(request.BookingID); err != nil {
		return err
	} else if userID != booking.UserID {
		return ErrNoPermission
	}

	// Soft delete entity
	return h.repository.DeleteBooking(request.BookingID)
}

func (h *Handler) GetBookingHistory(userID int, request apimodel.GetBookingHistoryRequest) (response []apimodel.GetBookingHistoryResponse, err error) {

	// Check request dates
	if request.StartDate.After(request.EndDate) {
		return response, ErrInvalidDates
	}

	// Get booking history from database with filters
	bookings, err := h.repository.GetBookingHistory(userID, request.MinPrice, request.MaxPrice, request.StartDate, request.EndDate)
	if err != nil {
		return
	}

	// Collect property identifiers
	var propertyIDs []int
	for _, b := range bookings {
		propertyIDs = append(propertyIDs, b.PropertyID)
	}

	// Get property data from database
	properties, err := h.repository.GetPropertiesUnscoped(propertyIDs)
	if err != nil {
		return
	}

	// Make map to compile data pieces
	propertyMap := make(map[int]dbmodel.Property)
	for _, p := range properties {
		propertyMap[p.ID] = p
	}

	// Create response model array
	for _, b := range bookings {
		days := 1
		if b.StartDate != b.EndDate {
			days = 1 + int(b.EndDate.Sub(b.StartDate).Hours()/24)
		}
		response = append(response, apimodel.GetBookingHistoryResponse{
			ID:          b.ID,
			Name:        propertyMap[b.PropertyID].Name,
			Zip:         propertyMap[b.PropertyID].Zip,
			Address:     propertyMap[b.PropertyID].Address,
			PricePerDay: b.PricePerDay,
			StartDate:   b.StartDate,
			EndDate:     b.EndDate,
			Days:        days,
			Price:       float64(days) * b.PricePerDay,
		})
	}
	return
}

func (h *Handler) Pay(userID int, request apimodel.PayRequest) (gatewayURL string, err error) {

	// Get booking data from database
	booking, err := h.repository.GetBooking(request.BookingID)
	if err != nil {
		return
	}

	// Check permission to pay
	if userID != booking.UserID {
		return gatewayURL, ErrNoPermission
	}

	// Check if the booking has been paid
	if booking.Paid {
		return gatewayURL, ErrPaid
	}

	// TODO: Create payment data struct and call api (sorry, but I didn't know how to mock this, I don't have much experience in it)

	// Return gateway URL
	return "test_gateway_url", nil
}
