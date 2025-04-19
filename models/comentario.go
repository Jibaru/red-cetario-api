package models

import "time"

type Comentario struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Descripcion string    `json:"descripcion"`
	IDCliente   uint      `json:"id_cliente"`
	IDReceta    uint      `json:"id_receta"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Comentario) TableName() string {
	return "comentarios"
}
