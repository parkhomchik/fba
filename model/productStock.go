package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type ProductStock struct {
	Id          uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Product     uuid.UUID `gorm:"type:uuid REFERENCES Product(Id)"`
	Point       uuid.UUID `gorm:"type:uuid REFERENCES Point(Id);"`
	Count       int64
	Changetimer time.Time
}

func (ProductStock) TableName() string {
	return "productstock"
}
