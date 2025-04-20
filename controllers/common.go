package controllers

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func GetLogger(c *gin.Context) *slog.Logger {
	if v, exists := c.Get("logger"); exists {
		if lg, ok := v.(*slog.Logger); ok {
			return lg
		}
	}
	return slog.Default()
}
