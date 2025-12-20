package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dronm/gobizapp/errs"
)

// TimeoutMiddleware sets a timeout for every request
func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a new context with timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		// Replace the context on the request
		c.Request = c.Request.WithContext(ctx)

		// Set up a channel to catch the request's completion
		done := make(chan struct{})
		go func() {
			c.Next()
			close(done)
		}()

		// Wait for either request completion or context timeout
		select {
		case <-done:
			// Request completed within timeout
		case <-ctx.Done():
			// Timeout exceeded
			code := errs.Timeout
			c.JSON(http.StatusRequestTimeout, gin.H{
				"error": errs.ErrorDescr(code), 
				"code": code,
			})
			c.Abort()
		}
	}
}
