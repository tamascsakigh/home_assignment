package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tamascsakigh/home_assignment/api/apimodel"
	"github.com/tamascsakigh/home_assignment/api/middleware"
	"github.com/tamascsakigh/home_assignment/handler"
	"net/http"
)

type Middleware interface {
	CORS(c *gin.Context)
	Authentication(c *gin.Context)
}

type Handler interface {
	GetFreeProperties(request apimodel.GetFreePropertiesRequest) ([]apimodel.GetFreePropertiesResponse, error)
	CreateBooking(userID int, request apimodel.CreateBookingRequest) error
	DeleteBooking(userID int, request apimodel.DeleteBookingRequest) error
	GetBookingHistory(userID int, request apimodel.GetBookingHistoryRequest) ([]apimodel.GetBookingHistoryResponse, error)
	Pay(userID int, request apimodel.PayRequest) (string, error)
}

type Router struct {
	middleware Middleware
	handler    Handler
}

func NewRouter(middleware Middleware, handler Handler) *Router {
	return &Router{middleware: middleware, handler: handler}
}

func (r *Router) InitApi() *gin.Engine {

	engine := gin.New()

	engine.Use(r.middleware.CORS, r.middleware.Authentication)

	engine.GET("/free", r.GetFreeProperties)
	engine.POST("/booking", r.PostBooking)
	engine.DELETE("/booking", r.DeleteBooking)
	engine.GET("/history", r.GetBookingHistory)
	engine.POST("/pay", r.PostPay)

	return engine
}

func (r *Router) GetFreeProperties(c *gin.Context) {

	var request apimodel.GetFreePropertiesRequest
	if err := c.BindJSON(&request); err != nil {
		return
	}
	if err := validator.New().Struct(request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := r.handler.GetFreeProperties(request)

	switch {
	case err == nil:
		c.JSON(http.StatusOK, response)
		return
	case errors.Is(err, handler.ErrInvalidDates):
		c.Status(http.StatusBadRequest)
		return
	default:
		c.Status(http.StatusInternalServerError)
	}
}

func (r *Router) PostBooking(c *gin.Context) {

	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var request apimodel.CreateBookingRequest
	if err = c.BindJSON(&request); err != nil {
		return
	}
	if err = validator.New().Struct(request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = r.handler.CreateBooking(userID, request)

	switch {
	case err == nil:
		c.Status(http.StatusOK)
		return
	case errors.Is(err, handler.ErrInvalidDates), errors.Is(err, handler.ErrDateConflict):
		c.Status(http.StatusBadRequest)
		return
	default:
		c.Status(http.StatusInternalServerError)
	}
}

func (r *Router) DeleteBooking(c *gin.Context) {

	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var request apimodel.DeleteBookingRequest
	if err = c.BindJSON(&request); err != nil {
		return
	}
	if err = validator.New().Struct(request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = r.handler.DeleteBooking(userID, request)

	switch {
	case err == nil:
		c.Status(http.StatusOK)
		return
	case errors.Is(err, handler.ErrNoPermission):
		c.Status(http.StatusForbidden)
		return
	default:
		c.Status(http.StatusInternalServerError)
	}
}

func (r *Router) GetBookingHistory(c *gin.Context) {

	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var request apimodel.GetBookingHistoryRequest
	if err = c.BindJSON(&request); err != nil {
		return
	}

	response, err := r.handler.GetBookingHistory(userID, request)

	switch {
	case err == nil:
		if request.Print {
			respJson, _ := json.MarshalIndent(response, "", "    ")
			fmt.Printf(string(respJson))
		}
		c.JSON(http.StatusOK, response)
		return
	case errors.Is(err, handler.ErrInvalidDates):
		c.Status(http.StatusBadRequest)
		return
	default:
		c.Status(http.StatusInternalServerError)
	}
}

func (r *Router) PostPay(c *gin.Context) {

	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var request apimodel.PayRequest
	if err = c.BindJSON(&request); err != nil {
		return
	}

	response, err := r.handler.Pay(userID, request)

	switch {
	case err == nil:
		c.Redirect(http.StatusOK, response)
		return
	case errors.Is(err, handler.ErrNoPermission):
		c.Status(http.StatusForbidden)
		return
	case errors.Is(err, handler.ErrPaid):
		c.Status(http.StatusBadRequest)
		return
	default:
		c.Status(http.StatusInternalServerError)
	}
}
