
package db

import "fba/model"
import "fmt"
import "github.com/satori/go.uuid"

//CREATE
func (dbm *DBManager) ActionTypeCreate(c model.ActionType) (err error) {
	if dbm.DB.NewRecord(&c) {
		err = dbm.DB.Create(&c).Error
		return err
	}
	return fmt.Errorf("%s", "запись уже существует")
}

//UPDATE
func (dbm *DBManager) ActionTypeUpdate(c model.ActionType) error {
	return dbm.DB.Save(&c).Error
}

//DELETE
func (dbm *DBManager) ActionTypeDelete(c model.ActionType) error {
	return dbm.DB.Delete(&c).Error
}

//GET
func (dbm *DBManager) ActionTypeGet(size, page int) (citys []model.ActionType, err error) {
	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&citys)
	return
}

//GET BY ID
func (dbm *DBManager) ActionTypeGetById(id uuid.UUID) (city model.ActionType, err error) {
	err = dbm.DB.Find(&city, id).Error
	return
}

//GET Count
func (dbm *DBManager) ActionTypeCount(size, page int) (count int, err error) {
	err = dbm.DB.Model(&model.ActionType{}).Count(&count).Error
	return
}
