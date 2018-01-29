package model

import uuid "github.com/satori/go.uuid"

type ProductActionDetail struct {
	Id             uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Product        uuid.UUID `gorm:"type:uuid REFERENCES Product(Id)"`
	Productsaction uuid.UUID `gorm:"type:uuid REFERENCES ProductAction(Id)"`
	Count          int64
	Pricebuy       float64
	Pricesell      float64
	Tax            uuid.UUID `gorm:"type:uuid REFERENCES Tax(Id);"`
	Reference      string    `gorm:"type:varchar;"`
	Partnumber     string    `gorm:"type:varchar;"`
	Actiontype     uuid.UUID `gorm:"type:uuid REFERENCES ActionType(Id)"`
	Paytype        uuid.UUID `gorm:"type:uuid REFERENCES PaymentType(Id)"`
}

func (ProductActionDetail) TableName() string {
	return "productactiondetail"
}
