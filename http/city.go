package http

import (
	"fmt"
	"strconv"

	"github.com/parkhomchik/fba/model"
	uuid "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
)

func (http *HttpManager) CityGET(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	citys, err := http.Manager.CityGet(size, page)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, citys)
}

func (http *HttpManager) CityGETByID(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	fmt.Println("search id =", id)
	if err != nil {
		c.JSON(400, err)
		return
	}
	citys, err := http.Manager.CityGetById(id)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, citys)
}

func (http *HttpManager) CityPOST(c *gin.Context) {
	var city model.City
	if c.Bind(&city) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	city, err := http.Manager.CityCreate(city)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, city)
	return
}

func (http *HttpManager) CityPUT(c *gin.Context) {
	var city model.City
	if c.Bind(&city) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	if err := http.Manager.CityUpdate(city); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, city)
	return
}

func (http *HttpManager) CityDELETE(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.JSON(400, err)
		return
	}

	city, err := http.Manager.CityGetById(id)

	if err := http.Manager.CityDelete(city); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "")
	return
}
