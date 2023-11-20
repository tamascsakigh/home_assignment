package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const UserID = "user_id"

type Middleware struct{}

func NewMiddleware() *Middleware { return &Middleware{} }

func (m *Middleware) CORS(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
}

func (m *Middleware) Authentication(c *gin.Context) {
	// mock authorization, add user_id to context
	c.Set(UserID, 1)
}

func GetUserID(c *gin.Context) (int, error) {
	value, exists := c.Get(UserID)
	userID, ok := value.(int)
	if !exists || !ok || userID < 1 {
		return 0, errors.New("invalid user_id")
	}
	return userID, nil
}
