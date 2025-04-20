package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"redcetarioapi/models"
)

func UpdateCliente(c *gin.Context) {
	logger := GetLogger(c)
	logger.InfoContext(c, "update client called", "path", c.Request.URL.Path)

	id := c.Param("id")
	var input struct {
		Nombre            string `json:"nombre"`
		ApePaterno        string `json:"ape_paterno"`
		ApeMaterno        string `json:"ape_materno"`
		CorreoElectronico string `json:"correo_electronico"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.ErrorContext(c, "update client failed", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "mensaje": err.Error()})
		return
	}
	DB.Model(&models.Cliente{}).Where("id = ?", id).Updates(models.Cliente{
		Nombre:            input.Nombre,
		ApePaterno:        input.ApePaterno,
		ApeMaterno:        input.ApeMaterno,
		CorreoElectronico: input.CorreoElectronico,
	})

	logger.InfoContext(c, "client updated", "client_id", id)

	c.JSON(http.StatusOK, gin.H{"ok": true})
}
