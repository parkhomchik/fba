package db

import "github.com/parkhomchik/fba/model"
import "fmt"
import "github.com/satori/go.uuid"

//CREATE
func (dbm *DBManager) CityCreate(c model.City) (city model.City, err error) {
	if dbm.DB.NewRecord(&c) {
		err = dbm.DB.Create(&c).Error
		return c, err
	}
	return c, fmt.Errorf("%s", "запись уже существует")
}

//UPDATE
func (dbm *DBManager) CityUpdate(c model.City) error {
	return dbm.DB.Save(&c).Error
}

//DELETE
func (dbm *DBManager) CityDelete(c model.City) error {
	return dbm.DB.Delete(&c).Error
}

//GET
func (dbm *DBManager) CityGet(size, page int) (citys []model.City, err error) {
	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&citys)
	return
}

//GET BY ID
func (dbm *DBManager) CityGetById(id uuid.UUID) (city model.City, err error) {
	err = dbm.DB.Where("id = ?", id).First(&city).Error
	return
}

//GET Count
func (dbm *DBManager) CityCount(size, page int) (count int, err error) {
	err = dbm.DB.Model(&model.City{}).Count(&count).Error
	return
}
