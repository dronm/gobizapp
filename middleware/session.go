package middleware

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	sess "github.com/dronm/session"
)

func extractCookieSessionID(c *gin.Context, cookieKey string) string {
	// cookie value if exists
	if vCookie, err := c.Cookie(cookieKey); err == nil {
		return vCookie
	}
	return ""
}

func extractHeaderSessionID(c *gin.Context) string {
	// cookie valuer
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return ""
	}
	return parts[1]
}

type SessionManager interface {
	SessionStart(string) (sess.Session, error)
	GetMaxLifeTime() int64
}

func SessionMiddleware(
	manager SessionManager,
	cookieKey string,
	isProduction bool,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/favicon.ico" {
			c.Next()
			return
		}
		sessID := extractCookieSessionID(c, cookieKey)
		if sessID == "" {
			sessID = extractHeaderSessionID(c) // for API calls without cookies
		}

		var sessionAlreadyExists bool
		if sessID != "" {
			sessionAlreadyExists = true
		}

		c.Set("session_loader", func() (sess.Session, error) {
			sess, err := manager.SessionStart(sessID)
			if err != nil {
				return nil, fmt.Errorf("manager.SessionStart(): %v" , err)
			}

			if !sessionAlreadyExists {
				now := time.Now()
				_ = sess.Put("time_created", now)
			}

			// set cookie
			c.SetCookie(
				cookieKey,
				sess.SessionID(),
				int(manager.GetMaxLifeTime()),
				"/",
				"",
				isProduction,
				true,
			)
			c.Set("session", sess)
			return sess, nil
		})

		c.Next()
	}
}
