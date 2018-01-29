package model

import uuid "github.com/satori/go.uuid"

type Product struct {
	Id       uuid.UUID  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name     string     `gorm:"type:varchar;"`
	Barcode  string     `gorm:"type:varchar;"`
	Category *uuid.UUID `gorm:"type:uuid REFERENCES Category(Id)"`
	Staff    uuid.UUID
}
