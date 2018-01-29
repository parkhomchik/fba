package model

import uuid "github.com/satori/go.uuid"

type Point struct {
	Id        uuid.UUID  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Staff     uuid.UUID  //идентификатор клиента в clients_Db
	City      *uuid.UUID `gorm:"type:uuid REFERENCES City(Id);default:null"`
	Address   string     `gorm:"type:varchar;"`
	Name      string     `gorm:"type:varchar;"`
	Isdeleted bool       `gorm:"type:bool;"` //Флаг удаления записи
	ClientID  *uuid.UUID
}
