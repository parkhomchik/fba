
package db

import "fba/model"
import "fmt"
import "github.com/satori/go.uuid"

//CREATE
func (dbm *DBManager) PointCreate(c model.Point) (err error) {
	if dbm.DB.NewRecord(&c) {
		err = dbm.DB.Create(&c).Error
		return err
	}
	return fmt.Errorf("%s", "запись уже существует")
}

//UPDATE
func (dbm *DBManager) PointUpdate(c model.Point) error {
	return dbm.DB.Save(&c).Error
}

//DELETE
func (dbm *DBManager) PointDelete(c model.Point) error {
	return dbm.DB.Delete(&c).Error
}

//GET
func (dbm *DBManager) PointGet(size, page int) (citys []model.Point, err error) {
	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&citys)
	return
}

//GET BY ID
func (dbm *DBManager) PointGetById(id uuid.UUID) (city model.Point, err error) {
	err = dbm.DB.Find(&city, id).Error
	return
}

//GET Count
func (dbm *DBManager) PointCount(size, page int) (count int, err error) {
	err = dbm.DB.Model(&model.Point{}).Count(&count).Error
	return
}
