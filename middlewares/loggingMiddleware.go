package middlewares

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Starting time request
		startTime := time.Now()

		// Processing request
		ctx.Next()

		// End Time request
		endTime := time.Now()

		// execution time
		latencyTime := endTime.Sub(startTime)
		latencyString := strconv.FormatInt(int64(latencyTime/time.Microsecond), 10) + " µs"

		log.WithFields(log.Fields{
			"METHOD":    ctx.Request.Method,
			"URI":       ctx.Request.RequestURI,
			"STATUS":    ctx.Writer.Status(),
			"LATENCY":   latencyString,
			"CLIENT_IP": ctx.ClientIP(),
		}).Info("HTTP REQUEST")

		ctx.Next()
	}
}
