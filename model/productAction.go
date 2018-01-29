package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type ProductAction struct {
	Id         uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Point      uuid.UUID `gorm:"type:uuid REFERENCES Point(Id);"`
	Session    uuid.UUID `gorm:"type:uuid REFERENCES WorkSession(Id);"`
	Count      int64
	Amount     float64
	Actiondate time.Time
}

func (ProductAction) TableName() string {
	return "productaction"
}
