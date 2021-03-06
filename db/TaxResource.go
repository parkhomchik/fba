package db

import "github.com/parkhomchik/fba/model"
import "fmt"
import "github.com/satori/go.uuid"

//CREATE
func (dbm *DBManager) TaxCreate(c model.Tax) (err error) {
	if dbm.DB.NewRecord(&c) {
		err = dbm.DB.Create(&c).Error
		return err
	}
	return fmt.Errorf("%s", "запись уже существует")
}

//UPDATE
func (dbm *DBManager) TaxUpdate(c model.Tax) error {
	return dbm.DB.Save(&c).Error
}

//DELETE
func (dbm *DBManager) TaxDelete(c model.Tax) error {
	return dbm.DB.Delete(&c).Error
}

//GET
func (dbm *DBManager) TaxGet(size, page int) (citys []model.Tax, err error) {
	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&citys)
	return
}

//GET BY ID
func (dbm *DBManager) TaxGetById(id uuid.UUID) (city model.Tax, err error) {
	err = dbm.DB.Find(&city, id).Error
	return
}

//GET Count
func (dbm *DBManager) TaxCount(size, page int) (count int, err error) {
	err = dbm.DB.Model(&model.Tax{}).Count(&count).Error
	return
}
