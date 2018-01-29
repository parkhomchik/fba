package db

import "github.com/parkhomchik/fba/model"
import "fmt"
import "github.com/satori/go.uuid"

//CREATE
func (dbm *DBManager) ProductCreate(c model.Product) (err error) {
	if dbm.DB.NewRecord(&c) {
		err = dbm.DB.Create(&c).Error
		return err
	}
	return fmt.Errorf("%s", "запись уже существует")
}

//UPDATE
func (dbm *DBManager) ProductUpdate(c model.Product) error {
	return dbm.DB.Save(&c).Error
}

//DELETE
func (dbm *DBManager) ProductDelete(c model.Product) error {
	return dbm.DB.Delete(&c).Error
}

//GET
func (dbm *DBManager) ProductGet(size, page int) (citys []model.Product, err error) {
	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&citys)
	return
}

//GET BY ID
func (dbm *DBManager) ProductGetById(id uuid.UUID) (city model.Product, err error) {
	err = dbm.DB.Find(&city, id).Error
	return
}

//GET Count
func (dbm *DBManager) ProductCount(size, page int) (count int, err error) {
	err = dbm.DB.Model(&model.Product{}).Count(&count).Error
	return
}
