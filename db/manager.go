package db

import (
	"github.com/parkhomchik/fba/model"

	"github.com/jinzhu/gorm"
)

//DBManager для связывания методов БД
type DBManager struct {
	DB *gorm.DB
}

//Init инициализация таблиц в БД
func (dbm *DBManager) Init() {
	dbm.DB.LogMode(true)
	dbm.DB.SingularTable(true)
	dbm.DB.AutoMigrate(&model.City{}, &model.Point{}, &model.Category{}, &model.Tax{}, &model.Role{}, &model.Staff{}, &model.Product{}, &model.ActionType{}, &model.PaymentType{}, &model.WorkSession{}, &model.ProductAction{}, &model.ProductActionDetail{}, &model.ProductStock{}, &model.StaffPoint{})
}
