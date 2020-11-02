package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Watson-Sei/watson-sei-official/api_v1/Config"
	"github.com/Watson-Sei/watson-sei-official/api_v1/Models"
	"github.com/Watson-Sei/watson-sei-official/api_v1/Routes"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var err error
var count = 0

func connection() *gorm.DB {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		log.Println("Not ready, Retry connecting...")
		time.Sleep(time.Second)
		count++
		log.Println(count)
		if count > 30 {
			panic(err)
		}
		return connection()
	}
	return Config.DB
}

func main() {
	gin.SetMode(gin.DebugMode)
	connection()
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status: ", err)
	}
	defer Config.DB.Close()
	// Migrate
	Config.DB.AutoMigrate(&Models.Article{}, &Models.User{})

	router := Routes.SetupRouter()
	// running
	router.Run(":8080")
}