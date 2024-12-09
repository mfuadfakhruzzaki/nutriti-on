// backend/middlewares/logger.go
package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
    logger := logrus.New()
    logger.SetFormatter(&logrus.JSONFormatter{})

    return func(c *gin.Context) {
        startTime := time.Now()

        // Process request
        c.Next()

        // Calculate latency
        latency := time.Since(startTime)

        // Get status
        status := c.Writer.Status()

        // Log entry
        entry := logger.WithFields(logrus.Fields{
            "status":     status,
            "method":     c.Request.Method,
            "path":       c.Request.URL.Path,
            "ip":         c.ClientIP(),
            "latency":    latency.String(),
            "user_agent": c.Request.UserAgent(),
        })

        if len(c.Errors) > 0 {
            entry.Error(c.Errors.String())
        } else {
            entry.Info("request processed")
        }
    }
}
