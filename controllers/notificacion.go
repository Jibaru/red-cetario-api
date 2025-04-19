package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"redcetarioapi/models"
)

func GetNotificaciones(c *gin.Context) {
	var notas []models.Notificacion
	DB.Find(&notas)
	c.JSON(http.StatusOK, gin.H{"ok": len(notas) > 0, "notificaciones": notas})
}

func UpdateNotificacion(c *gin.Context) {
	id := c.Param("id")
	now := time.Now()
	DB.Model(&models.Notificacion{}).Where("id = ?", id).Update("fecha_visto", now)
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func DeleteNotificacion(c *gin.Context) {
	id := c.Param("id")
	DB.Delete(&models.Notificacion{}, id)
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func DeleteNotificacionesByCliente(c *gin.Context) {
	id := c.Param("id")
	DB.Where("id_cliente = ?", id).Delete(&models.Notificacion{})
	c.JSON(http.StatusOK, gin.H{"ok": true})
}
