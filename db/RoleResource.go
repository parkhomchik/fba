package db

import "github.com/parkhomchik/fba/model"
import "fmt"
import "github.com/satori/go.uuid"

//CREATE
func (dbm *DBManager) RoleCreate(c model.Role) (err error) {
	if dbm.DB.NewRecord(&c) {
		err = dbm.DB.Create(&c).Error
		return err
	}
	return fmt.Errorf("%s", "запись уже существует")
}

//UPDATE
func (dbm *DBManager) RoleUpdate(c model.Role) error {
	return dbm.DB.Save(&c).Error
}

//DELETE
func (dbm *DBManager) RoleDelete(c model.Role) error {
	return dbm.DB.Delete(&c).Error
}

//GET
func (dbm *DBManager) RoleGet(size, page int) (citys []model.Role, err error) {
	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&citys)
	return
}

//GET BY ID
func (dbm *DBManager) RoleGetById(id uuid.UUID) (city model.Role, err error) {
	err = dbm.DB.Find(&city, id).Error
	return
}

//GET Count
func (dbm *DBManager) RoleCount(size, page int) (count int, err error) {
	err = dbm.DB.Model(&model.Role{}).Count(&count).Error
	return
}
