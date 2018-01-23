
package db

import "fba/model"
import "fmt"
import "github.com/satori/go.uuid"

//CREATE
func (dbm *DBManager) StaffCreate(c model.Staff) (err error) {
	if dbm.DB.NewRecord(&c) {
		err = dbm.DB.Create(&c).Error
		return err
	}
	return fmt.Errorf("%s", "запись уже существует")
}

//UPDATE
func (dbm *DBManager) StaffUpdate(c model.Staff) error {
	return dbm.DB.Save(&c).Error
}

//DELETE
func (dbm *DBManager) StaffDelete(c model.Staff) error {
	return dbm.DB.Delete(&c).Error
}

//GET
func (dbm *DBManager) StaffGet(size, page int) (citys []model.Staff, err error) {
	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&citys)
	return
}

//GET BY ID
func (dbm *DBManager) StaffGetById(id uuid.UUID) (city model.Staff, err error) {
	err = dbm.DB.Find(&city, id).Error
	return
}

//GET Count
func (dbm *DBManager) StaffCount(size, page int) (count int, err error) {
	err = dbm.DB.Model(&model.Staff{}).Count(&count).Error
	return
}
