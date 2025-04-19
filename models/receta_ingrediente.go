package models

type RecetaIngrediente struct {
	IDReceta      uint   `json:"id_receta"`
	IDIngrediente uint   `json:"id_ingrediente"`
	Cantidad      int    `json:"cantidad"`
	Unidad        string `json:"unidad"`
}

func (RecetaIngrediente) TableName() string {
	return "recetas_ingredientes"
}
