package http

import (
	"fmt"
	"strconv"

	"github.com/parkhomchik/fba/model"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type Paginator struct {
	Page int
	Size int
}

func (http *HttpManager) PointPOST(c *gin.Context) {
	var point model.Point
	if c.Bind(&point) != nil {
		c.JSON(400, "problem decoding body")
		return
	}

	if err := http.Manager.PointCreate(point); err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, point)
}

func (http *HttpManager) PointPUT(c *gin.Context) {
	var point model.Point
	if c.Bind(&point) != nil {
		c.JSON(400, "problem decoding body")
		return
	}

	if err := http.Manager.PointUpdate(point); err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, point)
}

func (http *HttpManager) PointDELETE(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.JSON(400, err)
		return
	}
	clientID := c.MustGet("clientID").(uuid.UUID)
	point, err := http.Manager.PointGetById(id, clientID)
	if err != nil {
		c.JSON(500, err)
		return
	}

	if err := http.Manager.PointDelete(point); err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, "")
}

func (http *HttpManager) PointGet(c *gin.Context) {

	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	clientID := c.MustGet("UserID").(uuid.UUID)

	fmt.Println("CLIENT ID:", clientID)

	points, err := http.Manager.PointGet(size, page, clientID)
	if err != nil {
		c.JSON(400, err)
	}

	c.JSON(200, points)
}

func (http *HttpManager) PointGetByID(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.JSON(400, err)
		return
	}
	clientID := c.MustGet("UserID").(uuid.UUID)
	points, err := http.Manager.PointGetById(id, clientID)
	if err != nil {
		c.JSON(400, err)
	}

	c.JSON(200, points)
}
