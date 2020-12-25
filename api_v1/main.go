package main

import (
	"github.com/Watson-Sei/watson-sei-official/api_v1/Config"
	"github.com/Watson-Sei/watson-sei-official/api_v1/Controllers"
	"github.com/Watson-Sei/watson-sei-official/api_v1/Models"
	"github.com/Watson-Sei/watson-sei-official/api_v1/Routes"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	gin.SetMode(gin.DebugMode)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: Config.DBUrl(Config.BuildDBConfig()),
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// v2になってからClose対処
	sqlDB, err := db.DB()
	defer sqlDB.Close()

	// Migrate
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Models.Article{}, &Models.User{}, &Models.Tag{})

	// Model *gorm.DB
	controller := Models.Model{Db: db}

	router := Routes.SetupRouter(Controllers.Controller{Model: controller})
	// Static Server
	router.Use(static.Serve("/assets", static.LocalFile("./assets", true)))
	// running
	router.Run(":8080")
}