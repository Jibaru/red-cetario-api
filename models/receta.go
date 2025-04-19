package models

import "time"

type Receta struct {
	ID                uint          `json:"id" gorm:"primaryKey"`
	Titulo            string        `json:"titulo"`
	Descripcion       string        `json:"descripcion"`
	TiempoPrep        int           `json:"tiempo_prep"`
	TiempoCoccion     int           `json:"tiempo_coccion"`
	URLImagen         string        `json:"url_imagen"`
	Tips              *string       `json:"tips,omitempty"`
	Calorias          string        `json:"calorias"`
	Dificultad        string        `json:"dificultad"`
	Cocina            string        `json:"cocina"`
	IDCliente         *uint         `json:"id_cliente,omitempty"`
	Cliente           Cliente       `json:"cliente,omitempty" gorm:"foreignKey:IDCliente"`
	Ingredientes      []Ingrediente `json:"ingredientes,omitempty" gorm:"many2many:recetas_ingredientes"`
	Materiales        []Material    `json:"materiales,omitempty" gorm:"many2many:recetas_materiales"`
	Pasos             []Paso        `json:"pasos,omitempty" gorm:"foreignKey:IDReceta"`
	ClientesFavoritos []Cliente     `json:"clientes_favoritos,omitempty" gorm:"many2many:recetas_favoritas"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
}

func (Receta) TableName() string {
	return "recetas"
}
