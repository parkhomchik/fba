package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Tax struct {
	Id        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name      string    `gorm:"type:varchar;"`
	Validfrom time.Time
	Value     float32
}
