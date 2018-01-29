package db

import "github.com/parkhomchik/fba/model"
import "fmt"
import "github.com/satori/go.uuid"

//CREATE
func (dbm *DBManager) ProductActionDetailCreate(c model.ProductActionDetail) (err error) {
	if dbm.DB.NewRecord(&c) {
		err = dbm.DB.Create(&c).Error
		return err
	}
	return fmt.Errorf("%s", "запись уже существует")
}

//UPDATE
func (dbm *DBManager) ProductActionDetailUpdate(c model.ProductActionDetail) error {
	return dbm.DB.Save(&c).Error
}

//DELETE
func (dbm *DBManager) ProductActionDetailDelete(c model.ProductActionDetail) error {
	return dbm.DB.Delete(&c).Error
}

//GET
func (dbm *DBManager) ProductActionDetailGet(size, page int) (citys []model.ProductActionDetail, err error) {
	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&citys)
	return
}

//GET BY ID
func (dbm *DBManager) ProductActionDetailGetById(id uuid.UUID) (city model.ProductActionDetail, err error) {
	err = dbm.DB.Find(&city, id).Error
	return
}

//GET Count
func (dbm *DBManager) ProductActionDetailCount(size, page int) (count int, err error) {
	err = dbm.DB.Model(&model.ProductActionDetail{}).Count(&count).Error
	return
}
