package db

import "github.com/parkhomchik/fba/model"
import "fmt"
import "github.com/satori/go.uuid"

//CREATE
func (dbm *DBManager) WorkSessionCreate(c model.WorkSession) (err error) {
	if dbm.DB.NewRecord(&c) {
		err = dbm.DB.Create(&c).Error
		return err
	}
	return fmt.Errorf("%s", "запись уже существует")
}

//UPDATE
func (dbm *DBManager) WorkSessionUpdate(c model.WorkSession) error {
	return dbm.DB.Save(&c).Error
}

//DELETE
func (dbm *DBManager) WorkSessionDelete(c model.WorkSession) error {
	return dbm.DB.Delete(&c).Error
}

//GET
func (dbm *DBManager) WorkSessionGet(size, page int) (citys []model.WorkSession, err error) {
	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&citys)
	return
}

//GET BY ID
func (dbm *DBManager) WorkSessionGetById(id uuid.UUID) (city model.WorkSession, err error) {
	err = dbm.DB.Find(&city, id).Error
	return
}

//GET Count
func (dbm *DBManager) WorkSessionCount(size, page int) (count int, err error) {
	err = dbm.DB.Model(&model.WorkSession{}).Count(&count).Error
	return
}
