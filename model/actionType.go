package model

import uuid "github.com/satori/go.uuid"

type ActionType struct {
	Id   uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name string    `gorm:"type:varchar;"`
}

func (ActionType) TableName() string {
	return "actiontype"
}
