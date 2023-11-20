package testdata

import (
	"github.com/tamascsakigh/home_assignment/api/apimodel"
	"github.com/tamascsakigh/home_assignment/database/dbmodel"
	"time"
)

var TestGetFreePropertiesPropertyModels = []dbmodel.Property{
	{
		ID:                 5,
		Name:               "B&B Hotel Budapest City",
		Zip:                1094,
		Address:            "Angyal u. 1-3",
		CurrentPricePerDay: 17295,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	},
	{
		ID:                 1,
		Name:               "Medos Hotel Budapest",
		Zip:                1061,
		Address:            "Jókai tér 9",
		CurrentPricePerDay: 18204,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	},
}

var TestGetFreePropertiesRequests = []apimodel.GetFreePropertiesRequest{
	{
		MinPrice:  10000,
		MaxPrice:  25000,
		StartDate: time.Date(2024, 1, 10, 0, 0, 0, 0, time.Now().Location()),
		EndDate:   time.Date(2024, 1, 15, 0, 0, 0, 0, time.Now().Location()),
	},
	{
		MinPrice:  10000,
		MaxPrice:  25000,
		StartDate: time.Date(2024, 1, 15, 0, 0, 0, 0, time.Now().Location()),
		EndDate:   time.Date(2024, 1, 10, 0, 0, 0, 0, time.Now().Location()),
	},
	{
		MinPrice:  10000,
		MaxPrice:  25000,
		StartDate: time.Date(2020, 1, 10, 0, 0, 0, 0, time.Now().Location()),
		EndDate:   time.Date(2020, 1, 15, 0, 0, 0, 0, time.Now().Location()),
	},
}

var TestGetFreePropertiesResponse = []apimodel.GetFreePropertiesResponse{
	{
		ID:          5,
		Name:        "B&B Hotel Budapest City",
		Zip:         1094,
		Address:     "Angyal u. 1-3",
		PricePerDay: 17295,
	},
	{
		ID:          1,
		Name:        "Medos Hotel Budapest",
		Zip:         1061,
		Address:     "Jókai tér 9",
		PricePerDay: 18204,
	},
}

var TestCreateBookingPropertyModels = []dbmodel.Property{
	{
		ID:                 1,
		Name:               "Medos Hotel Budapest",
		Zip:                1061,
		Address:            "Jókai tér 9",
		CurrentPricePerDay: 18204,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	},
}

var TestCreateBookingBookingModels = []dbmodel.Booking{
	{
		ID:          7,
		UserID:      3,
		PropertyID:  2,
		StartDate:   time.Date(2024, 1, 10, 0, 0, 0, 0, time.Now().Location()),
		EndDate:     time.Date(2024, 1, 19, 0, 0, 0, 0, time.Now().Location()),
		PricePerDay: 37546,
		Paid:        true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}

var TestCreateBookingRequests = []apimodel.CreateBookingRequest{
	{
		PropertyID: 1,
		StartDate:  time.Date(2024, 1, 10, 0, 0, 0, 0, time.Now().Location()),
		EndDate:    time.Date(2024, 1, 15, 0, 0, 0, 0, time.Now().Location()),
	},
	{
		PropertyID: 1,
		StartDate:  time.Date(2024, 1, 15, 0, 0, 0, 0, time.Now().Location()),
		EndDate:    time.Date(2024, 1, 10, 0, 0, 0, 0, time.Now().Location()),
	},
	{
		PropertyID: 1,
		StartDate:  time.Date(2020, 1, 10, 0, 0, 0, 0, time.Now().Location()),
		EndDate:    time.Date(2020, 1, 15, 0, 0, 0, 0, time.Now().Location()),
	},
	{
		PropertyID: 2,
		StartDate:  time.Date(2024, 1, 10, 0, 0, 0, 0, time.Now().Location()),
		EndDate:    time.Date(2024, 1, 15, 0, 0, 0, 0, time.Now().Location()),
	},
}

var TestDeleteBookingBookingModels = []dbmodel.Booking{
	{
		ID:          1,
		UserID:      1,
		PropertyID:  1,
		StartDate:   time.Date(2023, 2, 1, 0, 0, 0, 0, time.Now().Location()),
		EndDate:     time.Date(2023, 2, 7, 0, 0, 0, 0, time.Now().Location()),
		PricePerDay: 17465,
		Paid:        true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          6,
		UserID:      2,
		PropertyID:  1,
		StartDate:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.Now().Location()),
		EndDate:     time.Date(2024, 1, 7, 0, 0, 0, 0, time.Now().Location()),
		PricePerDay: 18204,
		Paid:        true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}

var TestDeleteBookingRequests = []apimodel.DeleteBookingRequest{
	{
		BookingID: 1,
	},
	{
		BookingID: 6,
	},
}

var TestGetBookingHistoryPropertyModels = []dbmodel.Property{
	{
		ID:                 3,
		Name:               "Up Hotel Budapest",
		Zip:                1067,
		Address:            "Csengery u. 31",
		CurrentPricePerDay: 26925,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	},
}

var TestGetBookingHistoryBookingModels = []dbmodel.Booking{
	{
		ID:          3,
		UserID:      1,
		PropertyID:  3,
		StartDate:   time.Date(2023, 6, 11, 0, 0, 0, 0, time.Now().Location()),
		EndDate:     time.Date(2023, 6, 12, 0, 0, 0, 0, time.Now().Location()),
		PricePerDay: 24566,
		Paid:        true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}

var TestGetBookingHistoryRequests = []apimodel.GetBookingHistoryRequest{
	{
		MinPrice:  10000,
		MaxPrice:  25000,
		StartDate: time.Date(2023, 3, 1, 0, 0, 0, 0, time.Now().Location()),
		EndDate:   time.Date(2023, 7, 1, 0, 0, 0, 0, time.Now().Location()),
	},
	{
		MinPrice:  10000,
		MaxPrice:  25000,
		StartDate: time.Date(2023, 7, 1, 0, 0, 0, 0, time.Now().Location()),
		EndDate:   time.Date(2023, 3, 1, 0, 0, 0, 0, time.Now().Location()),
	},
}

var TestGetBookingHistoryResponse = []apimodel.GetBookingHistoryResponse{
	{
		ID:          3,
		Name:        "Up Hotel Budapest",
		Zip:         1067,
		Address:     "Csengery u. 31",
		PricePerDay: 24566,
		StartDate:   time.Date(2023, 6, 11, 0, 0, 0, 0, time.Now().Location()),
		EndDate:     time.Date(2023, 6, 12, 0, 0, 0, 0, time.Now().Location()),
		Days:        2,
		Price:       49132,
	},
}

var TestPayRequests = []apimodel.PayRequest{
	{
		BookingID: 5,
	},
	{
		BookingID: 6,
	},
	{
		BookingID: 4,
	},
}

var TestPayBookingModels = []dbmodel.Booking{
	{
		ID:          5,
		UserID:      1,
		PropertyID:  5,
		StartDate:   time.Date(2023, 10, 3, 0, 0, 0, 0, time.Now().Location()),
		EndDate:     time.Date(2023, 10, 6, 0, 0, 0, 0, time.Now().Location()),
		PricePerDay: 18203,
		Paid:        false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          6,
		UserID:      2,
		PropertyID:  1,
		StartDate:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.Now().Location()),
		EndDate:     time.Date(2024, 1, 7, 0, 0, 0, 0, time.Now().Location()),
		PricePerDay: 18204,
		Paid:        true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          4,
		UserID:      1,
		PropertyID:  4,
		StartDate:   time.Date(2023, 8, 8, 0, 0, 0, 0, time.Now().Location()),
		EndDate:     time.Date(2023, 9, 2, 0, 0, 0, 0, time.Now().Location()),
		PricePerDay: 43499,
		Paid:        true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}
