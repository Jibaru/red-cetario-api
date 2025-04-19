package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"redcetarioapi/models"
)

func UpdateCliente(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		Nombre            string `json:"nombre"`
		ApePaterno        string `json:"ape_paterno"`
		ApeMaterno        string `json:"ape_materno"`
		CorreoElectronico string `json:"correo_electronico"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "mensaje": err.Error()})
		return
	}
	DB.Model(&models.Cliente{}).Where("id = ?", id).Updates(models.Cliente{
		Nombre:            input.Nombre,
		ApePaterno:        input.ApePaterno,
		ApeMaterno:        input.ApeMaterno,
		CorreoElectronico: input.CorreoElectronico,
	})
	c.JSON(http.StatusOK, gin.H{"ok": true})
}
