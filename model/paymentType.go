package model

import uuid "github.com/satori/go.uuid"

type PaymentType struct {
	Id   uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name string    `gorm:"type:varchar;"`
}

func (PaymentType) TableName() string {
	return "paymenttype"
}
