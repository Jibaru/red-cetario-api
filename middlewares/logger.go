package middlewares

import (
	"log"
	"log/slog"
	"redcetarioapi/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const LoggerKey string = "logger"

func Logger(url, appKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		batch := logger.NewSaaSBatchHandler(url, appKey)
		logger := slog.New(batch)

		c.Set(LoggerKey, logger)
		c.Set("correlation_id", uuid.NewString())

		c.Next()

		if err := batch.Flush(); err != nil {
			log.Printf("error flushing logs: %v\n", err)
		}
	}
}
