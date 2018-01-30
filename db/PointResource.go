package db

import "github.com/parkhomchik/fba/model"
import "fmt"
import "github.com/satori/go.uuid"

//CREATE
func (dbm *DBManager) PointCreate(c model.Point, ti model.TokenInfo) (err error) {
	/*
		-	если пользователя нет то пнх
		-	создаем запись в поинтах
		- 	отправляем инфу в oauth
		-	проставляем связь с клиентом из oauth
	*/
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
func (dbm *DBManager) PointGet(size, page int, clientID, userID uuid.UUID) (points []model.Point, err error) {
	if userID.String() != "00000000-0000-0000-0000-000000000000" {
		dbm.DB.Where("staff = ?", userID).Limit(size).Order("id asc").Offset((page - 1) * size).Find(&points)
	} else {
		dbm.DB.Where("client_id = ?", clientID).Limit(size).Order("id asc").Offset((page - 1) * size).Find(&points)
	}
	return
}

//GET BY ID
func (dbm *DBManager) PointGetById(id uuid.UUID, clientID, userID uuid.UUID) (point model.Point, err error) {
	if userID.String() != "00000000-0000-0000-0000-000000000000" {
		err = dbm.DB.Where("staff = ?", userID).Find(&point, id).Error
	} else {
		err = dbm.DB.Where("client_id = ? AND id = ?", clientID, id).First(&point).Error
	}
	return
}

//GET Count
func (dbm *DBManager) PointCount(size, page int, clientID, userID uuid.UUID) (count int, err error) {

	if userID.String() != "00000000-0000-0000-0000-000000000000" {
		err = dbm.DB.Model(&model.Point{}).Where("staff = ?", userID).Count(&count).Error
	} else {
		dbm.DB.Model(&model.Point{}).Where("client_id = ?", clientID).Limit(size).Order("id asc").Offset((page - 1) * size).Count(&count)
	}
	return
}
