package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"redcetarioapi/models"
)

func GetRecetas(c *gin.Context) {
	logger := GetLogger(c)
	logger.InfoContext(c, "get recipes called", "path", c.Request.URL.Path)

	var recetas []models.Receta
	DB.Preload("Cliente").Find(&recetas)

	logger.InfoContext(c, "recipes retrieved", "count", len(recetas))

	c.JSON(http.StatusOK, gin.H{"ok": len(recetas) > 0, "recetas": recetas})
}

func GetReceta(c *gin.Context) {
	logger := GetLogger(c)
	logger.InfoContext(c, "get recipe called", "path", c.Request.URL.Path)

	id := c.Param("id")
	var receta models.Receta
	if err := DB.Preload("Cliente").Preload("Ingredientes").Preload("Materiales").Preload("Pasos").Preload("ClientesFavoritos").Preload("Comentarios").First(&receta, id).Error; err != nil {
		logger.ErrorContext(c, "recipe not found", "error", err)
		c.JSON(http.StatusNotFound, gin.H{"ok": false, "mensaje": "Receta no encontrada"})
		return
	}

	logger.InfoContext(c, "recipe retrieved", "recipe_id", receta.ID)

	c.JSON(http.StatusOK, gin.H{"ok": true, "receta": receta})
}

func PostComentario(c *gin.Context) {
	logger := GetLogger(c)
	logger.InfoContext(c, "create comment called", "path", c.Request.URL.Path)

	id := c.Param("id")
	var input struct {
		Descripcion string `json:"descripcion"`
		IDCliente   uint   `json:"id_cliente"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.ErrorContext(c, "create comment failed", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "mensaje": err.Error()})
		return
	}
	coment := models.Comentario{Descripcion: input.Descripcion, IDReceta: parseUint(id), IDCliente: input.IDCliente}
	DB.Create(&coment)

	logger.InfoContext(c, "comment created", "comment_id", coment.ID)

	c.JSON(http.StatusOK, gin.H{"ok": true, "comentario": coment})
}
