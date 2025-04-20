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
	DB.Find(&recetas)

	logger.InfoContext(c, "recipes retrieved", "count", len(recetas))

	c.JSON(http.StatusOK, gin.H{"ok": len(recetas) > 0, "recetas": recetas})
}

func GetReceta(c *gin.Context) {
	logger := GetLogger(c)
	logger.InfoContext(c, "get recipe called", "path", c.Request.URL.Path)

	db := DB
	id := c.Param("id")

	// 1) Receta base
	var receta models.Receta
	if err := db.First(&receta, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"ok": false, "mensaje": "Receta no encontrada"})
		return
	}

	// 2) Totales
	var totalFav int64
	db.Model(&models.RecetaFavorita{}).
		Where("id_receta = ?", receta.ID).
		Count(&totalFav)

	var totalCom int64
	db.Model(&models.Comentario{}).
		Where("id_receta = ?", receta.ID).
		Count(&totalCom)

	// 3) Cliente
	var cliente models.Cliente
	if receta.IDCliente != nil {
		db.First(&cliente, *receta.IDCliente)
	}

	// 4) Comentarios
	var comentarios []models.Comentario
	db.
		Select("id, descripcion, created_at").
		Where("id_receta = ?", receta.ID).
		Find(&comentarios)

	// 5) Ingredientes + pivote
	type ingredienteResp struct {
		ID          uint   `json:"id"`
		Nombre      string `json:"nombre"`
		Descripcion string `json:"descripcion"`
		Pivot       struct {
			Cantidad int    `json:"cantidad"`
			Unidad   string `json:"unidad"`
		} `json:"pivot"`
	}
	var ings []models.RecetaIngrediente
	db.Where("id_receta = ?", receta.ID).Find(&ings)

	var ingredientes []ingredienteResp
	for _, piv := range ings {
		var ing models.Ingrediente
		db.First(&ing, piv.IDIngrediente)

		ir := ingredienteResp{
			ID:          ing.ID,
			Nombre:      ing.Nombre,
			Descripcion: ing.Descripcion,
		}
		ir.Pivot.Cantidad = piv.Cantidad
		ir.Pivot.Unidad = piv.Unidad
		ingredientes = append(ingredientes, ir)
	}

	// 6) Materiales
	var matsPiv []models.RecetaMaterial
	db.Where("id_receta = ?", receta.ID).Find(&matsPiv)

	var materiales []models.Material
	for _, mp := range matsPiv {
		var m models.Material
		db.First(&m, mp.IDMaterial)
		materiales = append(materiales, m)
	}

	// 7) Pasos
	var pasos []models.Paso
	db.
		Where("id_receta = ?", receta.ID).
		Order("numero_orden").
		Find(&pasos)

	// 8) Clientes favoritos (sólo ID)
	type clienteID struct {
		ID uint `json:"id"`
	}
	var favsPiv []models.RecetaFavorita
	db.Where("id_receta = ?", receta.ID).Find(&favsPiv)

	var clientesFav []clienteID
	for _, f := range favsPiv {
		clientesFav = append(clientesFav, clienteID{ID: f.IDCliente})
	}

	// 9) Respuesta final
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
		"receta": gin.H{
			"id":                 receta.ID,
			"titulo":             receta.Titulo,
			"descripcion":        receta.Descripcion,
			"tiempo_prep":        receta.TiempoPrep,
			"tiempo_coccion":     receta.TiempoCoccion,
			"url_imagen":         receta.URLImagen,
			"cocina":             receta.Cocina,
			"total_favoritos":    totalFav,
			"total_comentarios":  totalCom,
			"tips":               receta.Tips,
			"dificultad":         receta.Dificultad,
			"calorias":           receta.Calorias,
			"cliente":            cliente,
			"comentarios":        comentarios,
			"ingredientes":       ingredientes,
			"materiales":         materiales,
			"pasos":              pasos,
			"clientes_favoritos": clientesFav,
		},
	})
}

func PostComentario(c *gin.Context) {
	logger := GetLogger(c)
	logger.InfoContext(c, "create comment called", "path", c.Request.URL.Path)

	id := c.Param("id")
	var input struct {
		Descripcion string `json:"descripcion"`
		IDCliente   string `json:"id_cliente"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.ErrorContext(c, "create comment failed", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "mensaje": err.Error()})
		return
	}
	coment := models.Comentario{Descripcion: input.Descripcion, IDReceta: parseUint(id), IDCliente: parseUint(input.IDCliente)}
	DB.Create(&coment)

	logger.InfoContext(c, "comment created", "comment_id", coment.ID)

	c.JSON(http.StatusOK, gin.H{"ok": true, "comentario": coment})
}
