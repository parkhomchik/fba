package db

import "github.com/parkhomchik/fba/model"
import "fmt"
import "github.com/satori/go.uuid"

//CREATE
func (dbm *DBManager) PointCreate(c model.Point, ti model.TokenInfo) (newPoint model.Point, err error) {
	if ti.UserIsNull() {
		return
	}

	c.Staff, _ = ti.GetUserID()
	if dbm.DB.NewRecord(&c) {
		err = dbm.DB.Create(&c).Error
		return c, err
	}
	return c, fmt.Errorf("%s", "запись уже существует")
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
func (dbm *DBManager) PointGet(size, page int, ti model.TokenInfo) (points []model.Point, err error) {
	uid, err := ti.GetUserID()
	cid, err := ti.GetClientID()

	if err != nil {
		return
	}

	if !ti.UserIsNull() {
		dbm.DB.Where("staff = ?", uid).Limit(size).Order("id asc").Offset((page - 1) * size).Find(&points)
	} else {
		dbm.DB.Where("client_id = ?", cid).Limit(size).Order("id asc").Offset((page - 1) * size).Find(&points)
	}
	return
}

//GET BY ID
func (dbm *DBManager) PointGetById(id uuid.UUID, ti model.TokenInfo) (point model.Point, err error) {
	uid, err := ti.GetUserID()
	cid, err := ti.GetClientID()

	if err != nil {
		return
	}
	if !ti.UserIsNull() {
		err = dbm.DB.Where("staff = ?", uid).Find(&point, id).Error
	} else {
		err = dbm.DB.Where("client_id = ? AND id = ?", cid, id).First(&point).Error
	}
	return
}

//GET Count
func (dbm *DBManager) PointCount(size, page int, ti model.TokenInfo) (count int, err error) {
	uid, err := ti.GetUserID()
	if err != nil {
		return
	}
	if !ti.UserIsNull() {
		err = dbm.DB.Model(&model.Point{}).Where("staff = ?", uid).Count(&count).Error
	} else {
		return 0, fmt.Errorf("%s", "No USER")
	}
	return
}
