package models

type RecetaMaterial struct {
	IDReceta   uint `json:"id_receta"`
	IDMaterial uint `json:"id_material"`
}

func (RecetaMaterial) TableName() string {
	return "recetas_materiales"
}
