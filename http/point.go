package http

import (
	"fmt"
	"strconv"
	"time"

	"github.com/parkhomchik/fba/model"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"encoding/json"
)

//Client oauth2 client
type Client struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Secret    string
	Domain    string
	UserID    uuid.UUID
}

//PointPOST Create point
func (http *HttpManager) PointPOST(c *gin.Context) {
	var point model.Point
	if c.Bind(&point) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	tokenInfo := c.MustGet("TokenInfo").(model.TokenInfo)
	p, err := http.Manager.PointCreate(point, tokenInfo)
	if err != nil {
		c.JSON(500, err)
		return
	}

	body, status, err := http.Send("POST", "http://localhost:9096/connect/registrationclient", c.Request.Header.Get("Authorization"))
	if err != nil {
		c.JSON(500, err)
		return
	}
	if status != 200 {
		http.Manager.PointDelete(point)
		c.JSON(500, "no ouath")
		return
	}
	client := Client{}
	err = json.Unmarshal(body, &client)
	p.ClientID = &client.ID

	if err = http.Manager.PointUpdate(p); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, p)
}

func (http *HttpManager) PointClientInfo(c *gin.Context) {
	clientid := c.Param("pointid")
	fmt.Println(clientid)
	body, status, err := http.Send("GET", "http://localhost:9096/connect/clientinfo/"+clientid, c.Request.Header.Get("Authorization"))
	fmt.Println(status, body)
	client := Client{}
	err = json.Unmarshal(body, &client)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, client)
}

//PointPUT update point
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

//PointDELETE delete point
func (http *HttpManager) PointDELETE(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.JSON(400, err)
		return
	}
	tokenInfo := c.MustGet("TokenInfo").(model.TokenInfo)
	point, err := http.Manager.PointGetById(id, tokenInfo)
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

//PointGet get list point, parameters page(1) and size(10)
func (http *HttpManager) PointGet(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	tokenInfo := c.MustGet("TokenInfo").(model.TokenInfo)
	points, err := http.Manager.PointGet(size, page, tokenInfo)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, points)
}

//PointGetByID get point by id
func (http *HttpManager) PointGetByID(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.JSON(400, err)
		return
	}
	tokenInfo := c.MustGet("TokenInfo").(model.TokenInfo)
	points, err := http.Manager.PointGetById(id, tokenInfo)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, points)
}

//PointCount get count points
func (http *HttpManager) PointCount(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	tokenInfo := c.MustGet("TokenInfo").(model.TokenInfo)
	count, err := http.Manager.PointCount(size, page, tokenInfo)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, count)
}
