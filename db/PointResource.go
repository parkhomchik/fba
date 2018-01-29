package db

import "github.com/parkhomchik/fba/model"
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
func (dbm *DBManager) PointGet(size, page int, clientID, userID uuid.UUID) (points []model.Point, err error) {

	/*
		- ожидаем userID и clientID
		- если есть пользователь то отбираем по нему
		- если пользователя нет то отбираем по клиенту
	*/

	fmt.Println("USERID =", userID, "CLIENTID =", clientID)

	if userID.String() != "00000000-0000-0000-0000-000000000000" {
		dbm.DB.Where("staff = ?", userID).Limit(size).Order("id asc").Offset((page - 1) * size).Find(&points)
	} else {
		//var sf model.StaffPoint
		//dbm.DB.Where("point = ? AND staff = ?", clientID).First(&sf)
		dbm.DB.Where("client_id = ?", clientID).Limit(size).Order("id asc").Offset((page - 1) * size).Find(&points)
	}

	return
}

//GET BY ID
func (dbm *DBManager) PointGetById(id uuid.UUID, clientID uuid.UUID) (point model.Point, err error) {
	err = dbm.DB.Where("staff = ?", clientID).Find(&point, id).Error
	return
}

//GET Count
func (dbm *DBManager) PointCount(size, page int) (count int, err error) {
	err = dbm.DB.Model(&model.Point{}).Count(&count).Error
	return
}
