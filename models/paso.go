package models

import "time"

type Paso struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	NumeroOrden int       `json:"numero_orden"`
	Contenido   string    `json:"contenido"`
	IDReceta    uint      `json:"id_receta"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Paso) TableName() string {
	return "pasos"
}
