package main

import (
	"github.com/Watson-Sei/watson-sei-official/api_v1/Config"
	"github.com/Watson-Sei/watson-sei-official/api_v1/Models"
	"github.com/Watson-Sei/watson-sei-official/api_v1/Routes"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var err error

func main() {
	gin.SetMode(gin.DebugMode)
	db := Config.DBConnect()
	defer db.Close()
	// Migrate
	db.AutoMigrate(&Models.Article{}, &Models.User{})

	router := Routes.SetupRouter()
	// Static Server
	router.Use(static.Serve("/assets", static.LocalFile("./assets", true)))
	// running
	router.Run(":8080")
}