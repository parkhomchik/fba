package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type WorkSession struct {
	Id        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name      string    `gorm:"type:varchar;"`
	Startdate time.Time
	Enddate   time.Time
	Isactive  bool
	Point     uuid.UUID `gorm:"type:uuid REFERENCES Point(Id);"`
}

func (WorkSession) TableName() string {
	return "worksession"
}
