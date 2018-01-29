package model

import uuid "github.com/satori/go.uuid"

type StaffPoint struct {
	Id    uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Point uuid.UUID `gorm:"type:uuid REFERENCES Point(Id)"`
	Staff uuid.UUID `gorm:"type:uuid REFERENCES Staff(Id)"`
}

func (StaffPoint) TableName() string {
	return "staffpoint"
}
