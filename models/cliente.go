package models

import "time"

type Cliente struct {
	ID                uint      `json:"id" gorm:"primaryKey"`
	Nombre            string    `json:"nombre"`
	ApePaterno        string    `json:"ape_paterno"`
	ApeMaterno        string    `json:"ape_materno"`
	Contrasenia       string    `json:"-"`
	CorreoElectronico string    `json:"correo_electronico"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (Cliente) TableName() string {
	return "clientes"
}
