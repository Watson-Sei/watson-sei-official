package Config

import (
	"fmt"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB
var count = 0
var ACCESS_TOKEN_SECRETKEY = os.Getenv("ACCESS_TOKEN_SECRETKEY")
var REFRESH_TOKEN_SECRETKEY = os.Getenv("REFRESH_TOKEN_SECRETKEY")

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
