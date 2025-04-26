package controllers

import (
	"log"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func GetLogger(c *gin.Context) *slog.Logger {
	if v, exists := c.Get("logger"); exists {
		log.Println("logger found in context")

		if lg, ok := v.(*slog.Logger); ok {
			log.Println("logger retrieved from context")
			return lg
		}
	}
	log.Println("logger not found in context, using default logger")
	return slog.Default()
}
