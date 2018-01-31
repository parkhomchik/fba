package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	fbaHTTP "github.com/parkhomchik/fba/http"
	"github.com/parkhomchik/fba/model"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var httpManager fbaHTTP.HttpManager

func main() {
	settings, err := loadConfiguration("settings.json")
	db, err := getDb(settings)
	check(err)
	httpManager.Manager.DB = db
	httpManager.Manager.Init()
	r := gin.Default()
	r.Use(setCORSMiddleware())
	base := r.Group("/", auth())

	base.OPTIONS("", func(c *gin.Context) { c.Next() })

	base.GET("city", httpManager.CityGET)
	base.GET("city/:id", httpManager.CityGETByID)
	base.POST("city", httpManager.CityPOST)
	base.PUT("city", httpManager.CityPUT)
	base.DELETE("city/:id", httpManager.CityDELETE)

	base.GET("point", httpManager.PointGet)
	base.GET("point/:id", httpManager.PointGetByID)
	base.POST("pint", httpManager.PointPOST)
	base.PUT("point", httpManager.PointPUT)
	base.DELETE("point/:id", httpManager.PointDELETE)

	r.Run(":" + strconv.Itoa(settings.HttpPort))
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//id, _ := uuid.FromString("59059364-3dbf-43f5-a366-ac7d5f5d903a")
		//c.Set("UserID", id)
		//тут нужно получать/проверять токен и сохранять pointid в переменную
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			token := c.Request.Header.Get("Authorization")
			var tokenInfo model.TokenInfo
			client := &http.Client{}
			req, err := http.NewRequest("GET", "http://localhost:9096/oauth2/check", nil)
			req.Header.Add("Authorization", token)
			resp, err := client.Do(req)
			if err != nil {
				c.AbortWithError(400, err)
			}
			if resp.StatusCode == 200 {
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					c.AbortWithStatusJSON(400, model.NewError("Problem with client information (body)"))
				}
				if err := json.Unmarshal(body, &tokenInfo); err != nil {
					c.AbortWithStatusJSON(400, model.NewError("Problem with client information (parsing)"))
				}
				fmt.Println("TOKENINFO", tokenInfo)
				c.Set("TokenInfo", tokenInfo)
				c.Next()
			} else if resp.StatusCode == 401 {
				c.AbortWithStatus(401)
			} else {
				c.AbortWithStatusJSON(400, model.NewError("Problem with authorization"))
			}
		}
	}
}

func check(err error) {
	if err != nil {
		fmt.Println("ERROR: ", err)
		panic(err)
	}
}

func loadConfiguration(file string) (config model.Settings, err error) {
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return
}

func getDb(settings model.Settings) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%v port=%v user=%v dbname=%v sslmode=disable password=%v", settings.Host, settings.Port, settings.User, settings.DBName, settings.Password)
	db, err := gorm.Open("postgres", connectionString)
	db.LogMode(true)
	db.SingularTable(true)
	db.DB().SetMaxOpenConns(80)
	db.DB().SetMaxIdleConns(9)
	return db, err
}

func setCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "authorization, content-type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
		c.Next()
	}
}
