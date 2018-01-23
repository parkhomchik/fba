package http

import (
	"fba/model"

	"github.com/gin-gonic/gin"
)

func (http *HttpManager) CityPOST(c *gin.Context) {
	var city model.City
	if c.Bind(&city) != nil {
		c.JSON(400, "problem decoding body")
		return
	}

	http.Manager.CityCreate(city)
}
