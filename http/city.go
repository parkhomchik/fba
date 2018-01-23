package http

import (
	"fba/db"
	"fba/model"

	"github.com/gin-gonic/gin"
)

type HttpManager struct {
	Manager db.DBManager
}

func (http *HttpManager) CityPOST(c *gin.Context) {
	var city model.City
	if c.Bind(&city) != nil {
		c.JSON(400, "problem decoding body")
		return
	}

	http.Manager.CityCreate(city)
}
