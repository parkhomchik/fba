
package db

import "fba/model"
import "fmt"
import "github.com/satori/go.uuid"

//CREATE
func (dbm *DBManager) ProductStockCreate(c model.ProductStock) (err error) {
	if dbm.DB.NewRecord(&c) {
		err = dbm.DB.Create(&c).Error
		return err
	}
	return fmt.Errorf("%s", "запись уже существует")
}

//UPDATE
func (dbm *DBManager) ProductStockUpdate(c model.ProductStock) error {
	return dbm.DB.Save(&c).Error
}

//DELETE
func (dbm *DBManager) ProductStockDelete(c model.ProductStock) error {
	return dbm.DB.Delete(&c).Error
}

//GET
func (dbm *DBManager) ProductStockGet(size, page int) (citys []model.ProductStock, err error) {
	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&citys)
	return
}

//GET BY ID
func (dbm *DBManager) ProductStockGetById(id uuid.UUID) (city model.ProductStock, err error) {
	err = dbm.DB.Find(&city, id).Error
	return
}

//GET Count
func (dbm *DBManager) ProductStockCount(size, page int) (count int, err error) {
	err = dbm.DB.Model(&model.ProductStock{}).Count(&count).Error
	return
}
