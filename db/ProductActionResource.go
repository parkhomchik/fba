package db

import "github.com/parkhomchik/fba/model"
import "fmt"
import "github.com/satori/go.uuid"

//CREATE
func (dbm *DBManager) ProductActionCreate(c model.ProductAction) (err error) {
	if dbm.DB.NewRecord(&c) {
		err = dbm.DB.Create(&c).Error
		return err
	}
	return fmt.Errorf("%s", "запись уже существует")
}

//UPDATE
func (dbm *DBManager) ProductActionUpdate(c model.ProductAction) error {
	return dbm.DB.Save(&c).Error
}

//DELETE
func (dbm *DBManager) ProductActionDelete(c model.ProductAction) error {
	return dbm.DB.Delete(&c).Error
}

//GET
func (dbm *DBManager) ProductActionGet(size, page int) (citys []model.ProductAction, err error) {
	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&citys)
	return
}

//GET BY ID
func (dbm *DBManager) ProductActionGetById(id uuid.UUID) (city model.ProductAction, err error) {
	err = dbm.DB.Find(&city, id).Error
	return
}

//GET Count
func (dbm *DBManager) ProductActionCount(size, page int) (count int, err error) {
	err = dbm.DB.Model(&model.ProductAction{}).Count(&count).Error
	return
}
