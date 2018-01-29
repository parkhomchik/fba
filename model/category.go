package model

import uuid "github.com/satori/go.uuid"

type Category struct {
	Id       uuid.UUID  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name     string     `gorm:"type:varchar;"`
	ParentId *uuid.UUID `gorm:"type:uuid REFERENCES Category(Id);default:null"`
	Staff    uuid.UUID
}
