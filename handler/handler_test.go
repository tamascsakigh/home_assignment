package handler_test

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/tamascsakigh/home_assignment/api/apimodel"
	"github.com/tamascsakigh/home_assignment/database/dbmodel"
	"github.com/tamascsakigh/home_assignment/handler"
	"github.com/tamascsakigh/home_assignment/handler/testdata"
	"testing"
	"time"
)

func TestGetFreeProperties(t *testing.T) {

	tests := []struct {
		name             string
		request          apimodel.GetFreePropertiesRequest
		repoErr          error
		expectedResponse []apimodel.GetFreePropertiesResponse
		expectedError    error
	}{
		{"ok", testdata.TestGetFreePropertiesRequests[0], nil, testdata.TestGetFreePropertiesResponse, nil},
		{"dateErr1", testdata.TestGetFreePropertiesRequests[1], nil, []apimodel.GetFreePropertiesResponse(nil), handler.ErrInvalidDates},
		{"dateErr2", testdata.TestGetFreePropertiesRequests[2], nil, []apimodel.GetFreePropertiesResponse(nil), handler.ErrInvalidDates},
		{"repoErr", testdata.TestGetFreePropertiesRequests[0], errors.New("repo_err"), testdata.TestGetFreePropertiesResponse, errors.New("repo_err")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockRepository := MockRepository{}
			mockRepository.On("GetFreeProperties", float64(10000), float64(25000), time.Date(2024, 1, 10, 0, 0, 0, 0, time.Now().Location()), time.Date(2024, 1, 15, 0, 0, 0, 0, time.Now().Location())).Return(testdata.TestGetFreePropertiesPropertyModels, test.repoErr)

			mockHandler := handler.NewHandler(&mockRepository)

			response, err := mockHandler.GetFreeProperties(test.request)

			require.Equal(t, test.expectedResponse, response)
			require.Equal(t, test.expectedError, err)
		})
	}
}

func TestCreateBooking(t *testing.T) {

	tests := []struct {
		name          string
		request       apimodel.CreateBookingRequest
		repoErr1      error
		repoErr2      error
		repoErr3      error
		expectedError error
	}{
		{"ok", testdata.TestCreateBookingRequests[0], nil, nil, nil, nil},
		{"dateErr1", testdata.TestCreateBookingRequests[1], nil, nil, nil, handler.ErrInvalidDates},
		{"dateErr2", testdata.TestCreateBookingRequests[2], nil, nil, nil, handler.ErrInvalidDates},
		{"conflictErr", testdata.TestCreateBookingRequests[3], nil, nil, nil, handler.ErrDateConflict},
		{"repoErr1", testdata.TestCreateBookingRequests[0], errors.New("repo_err_1"), nil, nil, errors.New("repo_err_1")},
		{"repoErr2", testdata.TestCreateBookingRequests[0], nil, errors.New("repo_err_2"), nil, errors.New("repo_err_2")},
		{"repoErr3", testdata.TestCreateBookingRequests[0], nil, nil, errors.New("repo_err_3"), errors.New("repo_err_3")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockRepository := MockRepository{}
			mockRepository.On("GetPropertyBookings", 1, time.Date(2024, 1, 10, 0, 0, 0, 0, time.Now().Location()), time.Date(2024, 1, 15, 0, 0, 0, 0, time.Now().Location())).Return([]dbmodel.Booking{}, test.repoErr1)
			mockRepository.On("GetPropertyBookings", 2, time.Date(2024, 1, 10, 0, 0, 0, 0, time.Now().Location()), time.Date(2024, 1, 15, 0, 0, 0, 0, time.Now().Location())).Return(testdata.TestCreateBookingBookingModels, test.repoErr1)
			mockRepository.On("GetProperty", 1).Return(testdata.TestCreateBookingPropertyModels[0], test.repoErr2)
			mockRepository.On("CreateBooking", 1, 1, time.Date(2024, 1, 10, 0, 0, 0, 0, time.Now().Location()), time.Date(2024, 1, 15, 0, 0, 0, 0, time.Now().Location()), float64(18204)).Return(test.repoErr3)

			mockHandler := handler.NewHandler(&mockRepository)

			err := mockHandler.CreateBooking(1, test.request)

			require.Equal(t, test.expectedError, err)
		})
	}
}

func TestDeleteBooking(t *testing.T) {

	tests := []struct {
		name          string
		request       apimodel.DeleteBookingRequest
		repoErr1      error
		repoErr2      error
		expectedError error
	}{
		{"ok", testdata.TestDeleteBookingRequests[0], nil, nil, nil},
		{"permissionErr", testdata.TestDeleteBookingRequests[1], nil, nil, handler.ErrNoPermission},
		{"repoErr1", testdata.TestDeleteBookingRequests[0], errors.New("repo_err_1"), nil, errors.New("repo_err_1")},
		{"repoErr2", testdata.TestDeleteBookingRequests[0], nil, errors.New("repo_err_2"), errors.New("repo_err_2")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockRepository := MockRepository{}
			mockRepository.On("GetBooking", 1).Return(testdata.TestDeleteBookingBookingModels[0], test.repoErr1)
			mockRepository.On("GetBooking", 6).Return(testdata.TestDeleteBookingBookingModels[1], test.repoErr1)
			mockRepository.On("DeleteBooking", 1).Return(test.repoErr2)

			mockHandler := handler.NewHandler(&mockRepository)

			err := mockHandler.DeleteBooking(1, test.request)

			require.Equal(t, test.expectedError, err)
		})
	}
}

func TestGetBookingHistory(t *testing.T) {

	tests := []struct {
		name             string
		request          apimodel.GetBookingHistoryRequest
		repoErr1         error
		repoErr2         error
		expectedResponse []apimodel.GetBookingHistoryResponse
		expectedError    error
	}{
		{"ok", testdata.TestGetBookingHistoryRequests[0], nil, nil, testdata.TestGetBookingHistoryResponse, nil},
		{"dateErr", testdata.TestGetBookingHistoryRequests[1], nil, nil, []apimodel.GetBookingHistoryResponse(nil), handler.ErrInvalidDates},
		{"repoErr1", testdata.TestGetBookingHistoryRequests[0], errors.New("repo_err_1"), nil, []apimodel.GetBookingHistoryResponse(nil), errors.New("repo_err_1")},
		{"repoErr2", testdata.TestGetBookingHistoryRequests[0], nil, errors.New("repo_err_2"), []apimodel.GetBookingHistoryResponse(nil), errors.New("repo_err_2")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockRepository := MockRepository{}
			mockRepository.On("GetBookingHistory", 1, float64(10000), float64(25000), time.Date(2023, 3, 01, 0, 0, 0, 0, time.Now().Location()), time.Date(2023, 7, 01, 0, 0, 0, 0, time.Now().Location())).Return(testdata.TestGetBookingHistoryBookingModels, test.repoErr1)
			mockRepository.On("GetPropertiesUnscoped", []int{3}).Return(testdata.TestGetBookingHistoryPropertyModels, test.repoErr2)

			mockHandler := handler.NewHandler(&mockRepository)

			response, err := mockHandler.GetBookingHistory(1, test.request)

			require.Equal(t, test.expectedResponse, response)
			require.Equal(t, test.expectedError, err)
		})
	}
}

func TestPay(t *testing.T) {

	tests := []struct {
		name             string
		request          apimodel.PayRequest
		repoErr          error
		expectedResponse string
		expectedError    error
	}{
		{"ok", testdata.TestPayRequests[0], nil, "test_gateway_url", nil},
		{"permissionErr", testdata.TestPayRequests[1], nil, "", handler.ErrNoPermission},
		{"statusErr", testdata.TestPayRequests[2], nil, "", handler.ErrPaid},
		{"repoErr", testdata.TestPayRequests[0], errors.New("repo_err"), "", errors.New("repo_err")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockRepository := MockRepository{}
			mockRepository.On("GetBooking", 5).Return(testdata.TestPayBookingModels[0], test.repoErr)
			mockRepository.On("GetBooking", 6).Return(testdata.TestPayBookingModels[1], test.repoErr)
			mockRepository.On("GetBooking", 4).Return(testdata.TestPayBookingModels[2], test.repoErr)

			mockHandler := handler.NewHandler(&mockRepository)

			response, err := mockHandler.Pay(1, test.request)

			require.Equal(t, test.expectedResponse, response)
			require.Equal(t, test.expectedError, err)
		})
	}
}

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetFreeProperties(minPrice, maxPrice float64, startDate, endDate time.Time) ([]dbmodel.Property, error) {
	args := m.Called(minPrice, maxPrice, startDate, endDate)
	return args.Get(0).([]dbmodel.Property), args.Error(1)
}

func (m *MockRepository) GetPropertyBookings(ID int, startDate, endDate time.Time) ([]dbmodel.Booking, error) {
	args := m.Called(ID, startDate, endDate)
	return args.Get(0).([]dbmodel.Booking), args.Error(1)
}

func (m *MockRepository) GetBookingHistory(userID int, minPrice, maxPrice float64, startDate, endDate time.Time) ([]dbmodel.Booking, error) {
	args := m.Called(userID, minPrice, maxPrice, startDate, endDate)
	return args.Get(0).([]dbmodel.Booking), args.Error(1)
}

func (m *MockRepository) GetPropertiesUnscoped(IDs []int) ([]dbmodel.Property, error) {
	args := m.Called(IDs)
	return args.Get(0).([]dbmodel.Property), args.Error(1)
}

func (m *MockRepository) GetProperty(ID int) (dbmodel.Property, error) {
	args := m.Called(ID)
	return args.Get(0).(dbmodel.Property), args.Error(1)
}

func (m *MockRepository) GetBooking(ID int) (dbmodel.Booking, error) {
	args := m.Called(ID)
	return args.Get(0).(dbmodel.Booking), args.Error(1)
}

func (m *MockRepository) CreateBooking(userID, propertyID int, startDate, endDate time.Time, pricePerDay float64) error {
	args := m.Called(userID, propertyID, startDate, endDate, pricePerDay)
	return args.Error(0)
}

func (m *MockRepository) DeleteBooking(ID int) error {
	args := m.Called(ID)
	return args.Error(0)
}
