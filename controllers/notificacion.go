package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"redcetarioapi/models"
)

func GetNotificaciones(c *gin.Context) {
	logger := GetLogger(c)
	logger.InfoContext(c, "get notifications called", "path", c.Request.URL.Path)
	var notas []models.Notificacion
	DB.Find(&notas)

	logger.InfoContext(c, "notifications retrieved", "count", len(notas))

	c.JSON(http.StatusOK, gin.H{"ok": len(notas) > 0, "notificaciones": notas})
}

func UpdateNotificacion(c *gin.Context) {
	logger := GetLogger(c)
	logger.InfoContext(c, "update notification called", "path", c.Request.URL.Path)

	id := c.Param("id")
	now := time.Now()
	DB.Model(&models.Notificacion{}).Where("id = ?", id).Update("fecha_visto", now)

	logger.InfoContext(c, "notification updated", "notification_id", id)

	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func DeleteNotificacion(c *gin.Context) {
	logger := GetLogger(c)
	logger.InfoContext(c, "delete notification called", "path", c.Request.URL.Path)

	id := c.Param("id")
	DB.Delete(&models.Notificacion{}, id)

	logger.InfoContext(c, "notification deleted", "notification_id", id)

	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func DeleteNotificacionesByCliente(c *gin.Context) {
	logger := GetLogger(c)
	logger.InfoContext(c, "get notifications by client called", "path", c.Request.URL.Path)

	id := c.Param("id")
	DB.Where("id_cliente = ?", id).Delete(&models.Notificacion{})

	logger.InfoContext(c, "notifications deleted by client", "client_id", id)

	c.JSON(http.StatusOK, gin.H{"ok": true})
}
