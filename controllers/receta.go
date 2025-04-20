package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"redcetarioapi/models"
)

func GetRecetas(c *gin.Context) {
	var recetas []models.Receta
	DB.Preload("Cliente").Find(&recetas)
	c.JSON(http.StatusOK, gin.H{"ok": len(recetas) > 0, "recetas": recetas})
}

func GetReceta(c *gin.Context) {
	id := c.Param("id")
	var receta models.Receta
	if err := DB.Preload("Cliente").Preload("Ingredientes").Preload("Materiales").Preload("Pasos").Preload("ClientesFavoritos").Preload("Comentarios").First(&receta, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"ok": false, "mensaje": "Receta no encontrada"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true, "receta": receta})
}

func PostComentario(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		Descripcion string `json:"descripcion"`
		IDCliente   uint   `json:"id_cliente"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "mensaje": err.Error()})
		return
	}
	coment := models.Comentario{Descripcion: input.Descripcion, IDReceta: parseUint(id), IDCliente: input.IDCliente}
	DB.Create(&coment)
	c.JSON(http.StatusOK, gin.H{"ok": true, "comentario": coment})
}
