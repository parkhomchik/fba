package db

import "github.com/parkhomchik/fba/model"
import "fmt"
import "github.com/satori/go.uuid"

//CREATE
func (dbm *DBManager) PaymentTypeCreate(c model.PaymentType) (err error) {
	if dbm.DB.NewRecord(&c) {
		err = dbm.DB.Create(&c).Error
		return err
	}
	return fmt.Errorf("%s", "запись уже существует")
}

//UPDATE
func (dbm *DBManager) PaymentTypeUpdate(c model.PaymentType) error {
	return dbm.DB.Save(&c).Error
}

//DELETE
func (dbm *DBManager) PaymentTypeDelete(c model.PaymentType) error {
	return dbm.DB.Delete(&c).Error
}

//GET
func (dbm *DBManager) PaymentTypeGet(size, page int) (citys []model.PaymentType, err error) {
	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&citys)
	return
}

//GET BY ID
func (dbm *DBManager) PaymentTypeGetById(id uuid.UUID) (city model.PaymentType, err error) {
	err = dbm.DB.Find(&city, id).Error
	return
}

//GET Count
func (dbm *DBManager) PaymentTypeCount(size, page int) (count int, err error) {
	err = dbm.DB.Model(&model.PaymentType{}).Count(&count).Error
	return
}
