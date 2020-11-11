package Config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var count = 0
var SECRETKEY = os.Getenv("SECRETKEY")

// DBConfig represents db configuration
type DBConfig struct {
	Host 		string
	User 		string
	DBName 		string
	Password	string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:		"tcp(db)",
		User: 		os.Getenv("MYSQL_USER"),
		Password: 	os.Getenv("MYSQL_PASSWORD"),
		DBName: 	os.Getenv("MYSQL_DATABASE"),
	}
	return &dbConfig
}

func DBUrl(dbConfig *DBConfig) string {
	return fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.DBName,
	)
}

func DBConnect() *gorm.DB {
	db, err := gorm.Open("mysql", DBUrl(BuildDBConfig()))
	if err != nil {
		log.Println("Not ready, Retry connecting...")
		time.Sleep(time.Second)
		count++
		log.Println(count)
		if count > 30 {
			panic(err)
		}
		return DBConnect()
	}
	return db
}
