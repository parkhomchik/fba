
package db

import "fba/model"
import "fmt"
import "github.com/satori/go.uuid"

//CREATE
func (dbm *DBManager) CategoryCreate(c model.Category) (err error) {
	if dbm.DB.NewRecord(&c) {
		err = dbm.DB.Create(&c).Error
		return err
	}
	return fmt.Errorf("%s", "запись уже существует")
}

//UPDATE
func (dbm *DBManager) CategoryUpdate(c model.Category) error {
	return dbm.DB.Save(&c).Error
}

//DELETE
func (dbm *DBManager) CategoryDelete(c model.Category) error {
	return dbm.DB.Delete(&c).Error
}

//GET
func (dbm *DBManager) CategoryGet(size, page int) (citys []model.Category, err error) {
	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&citys)
	return
}

//GET BY ID
func (dbm *DBManager) CategoryGetById(id uuid.UUID) (city model.Category, err error) {
	err = dbm.DB.Find(&city, id).Error
	return
}

//GET Count
func (dbm *DBManager) CategoryCount(size, page int) (count int, err error) {
	err = dbm.DB.Model(&model.Category{}).Count(&count).Error
	return
}
