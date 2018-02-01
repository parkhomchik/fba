package http

import (
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

//PointPOST создать точку
func (http *HttpManager) PointPOST(c *gin.Context) {
	var point model.Point
	if c.Bind(&point) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	ti := c.MustGet("TokenInfo").(model.TokenInfo)
	p, err := http.Manager.PointCreate(point, ti)
	if err != nil {
		c.JSON(500, err)
		return
	}

	http.PointCreateOauthClient(p.Id, ti)
	c.JSON(200, p)
}

//PointCreateOauthClient добавить авторизацию для точки
func (http *HttpManager) PointCreateOauthClient(pid uuid.UUID, ti model.TokenInfo) (client Client, err error) {
	point, err := http.Manager.PointGetById(pid, ti)

	body, status, err := http.Send("POST", "http://localhost:9096/connect/registrationclient", ti.Token)
	if err != nil {
		return
	}
	if status == 200 {
		err = json.Unmarshal(body, &client)
		point.ClientID = &client.ID

		if err = http.Manager.PointUpdate(point); err != nil {
			return
		}
	}
	return
}

//PointGetOauthClient получаем информацию по клиенту из OAuth
func (http *HttpManager) PointGetOauthClient(c *gin.Context) {
	pid, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.JSON(400, err)
		return
	}
	ti := c.MustGet("TokenInfo").(model.TokenInfo)
	point, err := http.Manager.PointGetById(pid, ti)
	client := Client{}

	if point.ClientID == nil {
		client, err = http.PointCreateOauthClient(pid, ti)
		c.JSON(200, client)
		return
	}

	body, _, err := http.Send("GET", "http://localhost:9096/connect/clientinfo/"+point.ClientID.String(), c.Request.Header.Get("Authorization"))
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

	ti := c.MustGet("TokenInfo").(model.TokenInfo)
	p, err := http.Manager.PointGetById(point.Id, ti)
	uid, err := ti.GetUserID()

	if err != nil {
		c.JSON(500, err)
		return
	}

	if p.Staff != uid {
		c.JSON(404, "Staff не совпадает")
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
	ti := c.MustGet("TokenInfo").(model.TokenInfo)
	uid, err := ti.GetUserID()

	if err != nil {
		c.JSON(500, err)
		return
	}

	if point.Staff != uid {
		c.JSON(500, "Staff не совпадает")
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
