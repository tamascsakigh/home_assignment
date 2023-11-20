package repository

import (
	"github.com/tamascsakigh/home_assignment/database/dbmodel"
	"gorm.io/gorm"
	"time"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetFreeProperties(minPrice, maxPrice float64, startDate, endDate time.Time) (properties []dbmodel.Property, err error) {

	query := r.db.Where("properties.current_price_per_day >= ?", minPrice)
	if maxPrice > 0 {
		query = query.Where("properties.current_price_per_day <= ?", maxPrice)
	}

	var excludeIds []int
	if err = r.db.Model(&dbmodel.Booking{}).Where("bookings.start_date between ? and ? or bookings.end_date between ? and ? "+
		"or ? between bookings.start_date and bookings.end_date or ? between bookings.start_date and bookings.end_date",
		startDate, endDate, startDate, endDate, startDate, endDate).Pluck("bookings.property_id", &excludeIds).Error; err != nil {
		return
	}
	if len(excludeIds) > 0 {
		query = query.Where("properties.id not in (?)", excludeIds)
	}

	err = query.Order("properties.current_price_per_day").Find(&properties).Error
	return
}

func (r *Repository) GetPropertyBookings(ID int, startDate, endDate time.Time) (bookings []dbmodel.Booking, err error) {
	err = r.db.Where("bookings.property_id = ? and (bookings.start_date between ? and ? or bookings.end_date between ? and ? "+
		"or ? between bookings.start_date and bookings.end_date or ? between bookings.start_date and bookings.end_date)",
		ID, startDate, endDate, startDate, endDate, startDate, endDate).Find(&bookings).Error
	return
}

func (r *Repository) GetBookingHistory(userID int, minPrice, maxPrice float64, startDate, endDate time.Time) (bookings []dbmodel.Booking, err error) {

	query := r.db.Where("bookings.user_id = ? and bookings.price_per_day >= ?", userID, minPrice)
	if maxPrice > 0 {
		query = query.Where("bookings.price_per_day <= ?", maxPrice)
	}
	if !startDate.IsZero() {
		query = query.Where("bookings.start_date >= ?", startDate)
	}
	if !endDate.IsZero() {
		query = query.Where("bookings.end_date <= ?", endDate)
	}

	err = query.Find(&bookings).Error
	return
}

func (r *Repository) GetPropertiesUnscoped(IDs []int) (properties []dbmodel.Property, err error) {
	err = r.db.Unscoped().Where("properties.id in (?)", IDs).Find(&properties).Error
	return
}

func (r *Repository) GetProperty(ID int) (property dbmodel.Property, err error) {
	err = r.db.First(&property, ID).Error
	return
}

func (r *Repository) GetBooking(ID int) (booking dbmodel.Booking, err error) {
	err = r.db.First(&booking, ID).Error
	return
}

func (r *Repository) CreateBooking(userID, propertyID int, startDate, endDate time.Time, pricePerDay float64) error {
	return r.db.Create(&dbmodel.Booking{
		UserID:      userID,
		PropertyID:  propertyID,
		StartDate:   startDate,
		EndDate:     endDate,
		PricePerDay: pricePerDay,
	}).Error
}

func (r *Repository) DeleteBooking(ID int) error {
	return r.db.Delete(&dbmodel.Booking{ID: ID}).Error
}
