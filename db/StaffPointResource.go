package db

import "github.com/parkhomchik/fba/model"
import "fmt"
import "github.com/satori/go.uuid"

//CREATE
func (dbm *DBManager) StaffPointCreate(c model.StaffPoint) (err error) {
	if dbm.DB.NewRecord(&c) {
		err = dbm.DB.Create(&c).Error
		return err
	}
	return fmt.Errorf("%s", "запись уже существует")
}

//UPDATE
func (dbm *DBManager) StaffPointUpdate(c model.StaffPoint) error {
	return dbm.DB.Save(&c).Error
}

//DELETE
func (dbm *DBManager) StaffPointDelete(c model.StaffPoint) error {
	return dbm.DB.Delete(&c).Error
}

//GET
func (dbm *DBManager) StaffPointGet(size, page int) (citys []model.StaffPoint, err error) {
	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&citys)
	return
}

//GET BY ID
func (dbm *DBManager) StaffPointGetById(id uuid.UUID) (city model.StaffPoint, err error) {
	err = dbm.DB.Find(&city, id).Error
	return
}

//GET Count
func (dbm *DBManager) StaffPointCount(size, page int) (count int, err error) {
	err = dbm.DB.Model(&model.StaffPoint{}).Count(&count).Error
	return
}
