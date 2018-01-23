package http

import (
	"fba/model"
	"strconv"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func (http *HttpManager) PointPOST(c *gin.Context) {
	var point model.Point
	if c.Bind(&point) != nil {
		c.JSON(400, "problem decoding body")
		return
	}

	http.Manager.PointCreate(point)
}

func (http *HttpManager) PointGet(c *gin.Context) {
	var points []model.Point
	size, err := strconv.Atoi(c.Param("size"))
	page, err := strconv.Atoi(c.Param("page"))
	clientID := c.MustGet("clientID").(uuid.UUID)
	points, err = http.Manager.PointGet(size, page, clientID)
	if err != nil {
		c.JSON(400, err)
	}

	c.JSON(200, points)
}

func (http *HttpManager) PointGetByID(c *gin.Context) {
	var points []model.Point
	size, err := strconv.Atoi(c.Param("size"))
	page, err := strconv.Atoi(c.Param("page"))
	clientID := c.MustGet("clientID").(uuid.UUID)
	points, err = http.Manager.PointGet(size, page, clientID)
	if err != nil {
		c.JSON(400, err)
	}

	c.JSON(200, points)
}
