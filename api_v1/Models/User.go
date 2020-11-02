package Models

import (
	"github.com/Watson-Sei/watson-sei-official/api_v1/Config"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username string, password string) (err error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	// insert処理
	if err := Config.DB.Create(&User{username,string(hash)}).Error; err != nil {
		return err
	}
	return nil
}