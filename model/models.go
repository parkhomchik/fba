package model

import (
	"time"

	"github.com/satori/go.uuid"
)

type City struct {
	Id   uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name string    `gorm:"type:varchar;"`
}

type Point struct {
	Id        uuid.UUID  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Client    uuid.UUID  //идентификатор клиента в clients_Db
	City      *uuid.UUID `gorm:"type:uuid REFERENCES City(Id);default:null"`
	Address   string     `gorm:"type:varchar;"`
	Name      string     `gorm:"type:varchar;"`
	Isdeleted bool       `gorm:"type:bool;"` //Флаг удаления записи
}

type Category struct {
	Id       uuid.UUID  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name     string     `gorm:"type:varchar;"`
	ParentId *uuid.UUID `gorm:"type:uuid REFERENCES Category(Id);default:null"`
	Client   uuid.UUID
}

type Tax struct {
	Id        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name      string    `gorm:"type:varchar;"`
	Validfrom time.Time
	Value     float32
}

type Role struct {
	Id   uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	name string    `gorm:"type:varchar;"`
}

type Staff struct {
	Id       uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name     string    `gorm:"type:varchar;"`
	Password string    `gorm:"type:varchar;"`
	Role     uuid.UUID `gorm:"type:uuid REFERENCES Role(Id)"`
	Visible  bool      `gorm:"type:bool;"`
}

type StaffPoint struct {
	Id    uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Point uuid.UUID `gorm:"type:uuid REFERENCES Point(Id)"`
	Staff uuid.UUID `gorm:"type:uuid REFERENCES Staff(Id)"`
}

func (StaffPoint) TableName() string {
	return "staffpoint"
}

type Product struct {
	Id       uuid.UUID  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name     string     `gorm:"type:varchar;"`
	Barcode  string     `gorm:"type:varchar;"`
	Category *uuid.UUID `gorm:"type:uuid REFERENCES Category(Id)"`
	Client   uuid.UUID
}

type ActionType struct {
	Id   uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name string    `gorm:"type:varchar;"`
}

func (ActionType) TableName() string {
	return "actiontype"
}

type PaymentType struct {
	Id   uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name string    `gorm:"type:varchar;"`
}

func (PaymentType) TableName() string {
	return "paymenttype"
}

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
