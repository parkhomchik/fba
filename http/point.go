package http

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/parkhomchik/fba/model"

	h "net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"encoding/json"
)

type Client struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Secret    string
	Domain    string
	UserID    uuid.UUID
}

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

	req, err := h.NewRequest("POST", "http://localhost:9096/connect/registrationclient", nil)
	req.Header.Add("Authorization", c.Request.Header.Get("Authorization"))
	cl := h.Client{}
	resp, err := cl.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("http://localhost:9096/connect/registrationclient status =", resp.Status)

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	client := Client{}
	err = json.Unmarshal(bodyBytes, &client)

	fmt.Println(client)
	p.ClientID = &client.ID

	if err = http.Manager.PointUpdate(p); err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, p)
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
	tokenInfo := c.MustGet("TokenInfo").(model.TokenInfo)
	UserID, _ := tokenInfo.GetUserID()
	ClientID, _ := tokenInfo.GetClientID()
	point, err := http.Manager.PointGetById(id, ClientID, UserID)
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
	//var tokenInfo model.TokenInfo
	tokenInfo := c.MustGet("TokenInfo").(model.TokenInfo)
	UserID, _ := tokenInfo.GetUserID()
	ClientID, _ := tokenInfo.GetClientID()

	points, err := http.Manager.PointGet(size, page, ClientID, UserID)
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
	tokenInfo := c.MustGet("TokenInfo").(model.TokenInfo)
	UserID, _ := tokenInfo.GetUserID()
	ClientID, _ := tokenInfo.GetClientID()

	points, err := http.Manager.PointGetById(id, ClientID, UserID)

	if err != nil {
		c.JSON(400, err)
	}

	c.JSON(200, points)
}

func (http *HttpManager) PointCount(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	//var tokenInfo model.TokenInfo
	tokenInfo := c.MustGet("TokenInfo").(model.TokenInfo)
	UserID, _ := tokenInfo.GetUserID()
	ClientID, _ := tokenInfo.GetClientID()

	count, err := http.Manager.PointCount(size, page, ClientID, UserID)
	if err != nil {
		c.JSON(400, err)
	}

	c.JSON(200, count)
}
