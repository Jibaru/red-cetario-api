package models

import "time"

type Material struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Nombre      string    `json:"nombre"`
	Descripcion string    `json:"descripcion"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Material) TableName() string {
	return "materiales"
}
