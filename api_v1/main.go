package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var err error
var count = 0

func Initdb() (*gorm.DB, error) {
	DBMS := "mysql"
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := "tcp(db)"
	DBNAME := "docker_db"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		log.Println("Not ready, Retry connecting...")
		time.Sleep(time.Second)
		count++
		log.Println(count)
		if count > 30 {
			panic(err)
		}
		return Initdb()
	}
	return db, nil
}

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	Initdb()
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	router.Run(":8080")
}