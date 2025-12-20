package middleware

import (
	"net/http"
	
	"github.com/dronm/gobizapp/logger"

	
	"github.com/gin-gonic/gin"
)

func CorsMiddleware(baseURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		
		logger.Logger.Debugf("CorsMiddleware: origin: %s, baseURL: %s, IsSame: %v", origin, baseURL, origin==baseURL)

		if origin == baseURL {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)//
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, Origin, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		// Handle preflight requests
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
	}
}
