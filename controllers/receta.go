package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"redcetarioapi/models"
)

func GetRecetas(c *gin.Context) {
	db := DB
	logger := GetLogger(c)
	logger.InfoContext(c, "get recipes called", "path", c.Request.URL.Path)

	// 1) Obtener todas las recetas
	var recetas []models.Receta
	db.Find(&recetas)

	if len(recetas) == 0 {
		c.JSON(http.StatusOK, gin.H{"ok": false, "recetas": []any{}})
		return
	}

	// 2) Obtener favoritos por receta (map)
	type favCount struct {
		IDReceta uint
		Total    int
	}
	var favs []favCount
	db.
		Table("recetas_favoritas").
		Select("id_receta, COUNT(*) as total").
		Group("id_receta").
		Scan(&favs)

	favMap := make(map[uint]int)
	for _, f := range favs {
		favMap[f.IDReceta] = f.Total
	}

	// 3) Obtener todos los clientes relacionados por ID
	var clienteIDs []uint
	for _, r := range recetas {
		if r.IDCliente != nil {
			clienteIDs = append(clienteIDs, *r.IDCliente)
		}
	}

	var clientes []models.Cliente
	if len(clienteIDs) > 0 {
		db.Where("id IN ?", clienteIDs).Find(&clientes)
	}

	clienteMap := make(map[uint]models.Cliente)
	for _, c := range clientes {
		clienteMap[c.ID] = c
	}

	// 4) Armar respuesta
	var respuesta []gin.H
	for _, r := range recetas {
		clienteData := gin.H{}
		if r.IDCliente != nil {
			if cliente, ok := clienteMap[*r.IDCliente]; ok {
				clienteData = gin.H{
					"id":                 cliente.ID,
					"nombre":             cliente.Nombre,
					"ape_paterno":        cliente.ApePaterno,
					"ape_materno":        cliente.ApeMaterno,
					"correo_electronico": cliente.CorreoElectronico,
				}
			}
		}

		respuesta = append(respuesta, gin.H{
			"id":              r.ID,
			"titulo":          r.Titulo,
			"url_imagen":      r.URLImagen,
			"tiempo_prep":     r.TiempoPrep,
			"tiempo_coccion":  r.TiempoCoccion,
			"cocina":          r.Cocina,
			"total_favoritos": favMap[r.ID],
			"cliente":         clienteData,
		})
	}

	logger.InfoContext(c, "recipes retrieved", "count", len(respuesta))
	c.JSON(http.StatusOK, gin.H{"ok": true, "recetas": respuesta})
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

	// 2) Cliente
	var cliente models.Cliente
	if receta.IDCliente != nil {
		db.First(&cliente, *receta.IDCliente)
	}

	// 3) Comentarios
	var comentarios []models.Comentario
	db.Select("id, descripcion, created_at").
		Where("id_receta = ?", receta.ID).
		Find(&comentarios)

	// 4) Pasos
	var pasos []models.Paso
	db.Where("id_receta = ?", receta.ID).Order("numero_orden").Find(&pasos)

	// 5) Clientes favoritos
	type clienteID struct {
		ID uint `json:"id"`
	}
	var favsPiv []models.RecetaFavorita
	db.Where("id_receta = ?", receta.ID).Find(&favsPiv)

	var clientesFav []clienteID
	for _, f := range favsPiv {
		clientesFav = append(clientesFav, clienteID{ID: f.IDCliente})
	}

	// 6) Ingredientes + pivote
	type ingredienteResp struct {
		ID          uint   `json:"id"`
		Nombre      string `json:"nombre"`
		Descripcion string `json:"descripcion"`
		Pivot       struct {
			Cantidad int    `json:"cantidad"`
			Unidad   string `json:"unidad"`
		} `json:"pivot"`
	}

	var recetasIngs []models.RecetaIngrediente
	db.Where("id_receta = ?", receta.ID).Find(&recetasIngs)

	var ingIDs []uint
	for _, r := range recetasIngs {
		ingIDs = append(ingIDs, r.IDIngrediente)
	}

	var ingredientesDB []models.Ingrediente
	if len(ingIDs) > 0 {
		db.Where("id IN ?", ingIDs).Find(&ingredientesDB)
	}

	ingMap := make(map[uint]models.Ingrediente)
	for _, ing := range ingredientesDB {
		ingMap[ing.ID] = ing
	}

	var ingredientes []ingredienteResp
	for _, r := range recetasIngs {
		ing := ingMap[r.IDIngrediente]
		item := ingredienteResp{
			ID:          ing.ID,
			Nombre:      ing.Nombre,
			Descripcion: ing.Descripcion,
		}
		item.Pivot.Cantidad = r.Cantidad
		item.Pivot.Unidad = r.Unidad
		ingredientes = append(ingredientes, item)
	}

	// 7) Materiales
	var recetasMats []models.RecetaMaterial
	db.Where("id_receta = ?", receta.ID).Find(&recetasMats)

	var matIDs []uint
	for _, m := range recetasMats {
		matIDs = append(matIDs, m.IDMaterial)
	}

	var materiales []models.Material
	if len(matIDs) > 0 {
		db.Where("id IN ?", matIDs).Find(&materiales)
	}

	// 8) Totales
	totalFav := len(favsPiv)
	totalCom := len(comentarios)

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
