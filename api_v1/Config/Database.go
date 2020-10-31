package Config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var count = 0

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

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.DBName,
	)
}

//func Initdb() *gorm.DB {
//	DBMS := "mysql"
//	USER := os.Getenv("MYSQL_USER")
//	PASS := os.Getenv("MYSQL_PASSWORD")
//	PROTOCOL := "tcp(db)"
//	DBNAME := os.Getenv("MYSQL_DATABASE")
//	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
//	db, err := gorm.Open(DBMS, CONNECT)
//	if err != nil {
//		log.Println("Not ready, Retry Connecting...")
//		time.Sleep(time.Second)
//		count++
//		log.Println(count)
//		if count > 30 {
//			panic(err)
//		}
//		return Initdb()
//	}
//	return db
//}