package models

import "time"

type Notificacion struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Titulo      string     `json:"titulo"`
	Descripcion string     `json:"descripcion"`
	FechaEnvio  time.Time  `json:"fecha_envio"`
	FechaVisto  *time.Time `json:"fecha_visto,omitempty"`
	IDCliente   uint       `json:"id_cliente"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func (Notificacion) TableName() string {
	return "notificaciones"
}
