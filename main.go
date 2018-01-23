package main

import (
	"encoding/json"
	"fba/db"
	"fba/http"
	"fba/model"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

var dbm db.DBManager
var httpManager http.HttpManager

func main() {
	settings, err := loadConfiguration("settings.json")
	db, err := getDb(settings)

	check(err)

	dbm.DB = db
	httpManager.Manager = dbm

	r := gin.Default()
	base := r.Group("/", auth(), DBManager())

	//base.POST("city", http.CityPOST)
	//base.PUT("city/:id", http.CityPOST)
	//base.DELETE("city/:id", http.CityPOST)
	//base.GET("city/:size:page", http.CityPOST)
	base.GET("city/:id", httpManager.CityPOST)

}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//тут нужно получать/проверять токен и сохранять pointid в переменную
		c.Next()
		return
	}
}

func check(err error) {
	if err != nil {
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

func getDb(settings model.Settings) (gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%v port=%v user=%v dbname=%v sslmode=disable password=%v", settings.Host, settings.Port, settings.User, settings.DBName, settings.Password)
	db, err := gorm.Open("postgres", connectionString)
	db.LogMode(false)
	db.SingularTable(true)
	db.DB().SetMaxOpenConns(80)
	db.DB().SetMaxIdleConns(9)
	return *db, err
}

func DBManager() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("dataBase", dbm)
	}
}
