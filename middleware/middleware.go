package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var limiters = make(map[string]*rate.Limiter)

// Limits the number of requests allowed per client IP
func RateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		if _, exists := limiters[ip]; !exists {
			limiters[ip] = rate.NewLimiter(1, 5) // 1 request per second with a burst capacity of 5
		}
		limiter := limiters[ip]

		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// Handles errors and provide consistent response format
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
}
