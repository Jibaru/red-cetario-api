package models

type RecetaFavorita struct {
	IDReceta  uint `json:"id_receta"`
	IDCliente uint `json:"id_cliente"`
}

func (RecetaFavorita) TableName() string {
	return "recetas_favoritas"
}
