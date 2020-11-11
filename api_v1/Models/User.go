package Models

import (
	"time"

	"github.com/Watson-Sei/watson-sei-official/api_v1/Config"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username string, password string) (err error) {
	db := Config.DBConnect()
	defer db.Close()
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	// insert処理
	if err := db.Create(&User{username,string(hash)}).Error; err != nil {
		return err
	}
	return nil
}

func CreateJWTToken(username string) string {
	/*
	アルゴリズム指定
	 */
	// header
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	// claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 4).Unix()

	// Electronic signature
	tokenString, err := token.SignedString([]byte(Config.SECRETKEY))
	if err == nil {
		return tokenString
	} else {
		return "token生成に失敗しました"
	}
}


// Redis JWT Token Black List Register
func BlackListSet(exp int64, token string) error {
	conn := Config.RedisConnection()
	defer conn.Close()

	// 残り時間
	nowTime := time.Now()
	expTime := time.Unix(exp, 0)

	// 残り時間秒数
	timeLeft := expTime.Sub(nowTime).Seconds()

	// Redis DBに追加
	_, err := conn.Do("SET", token, string(exp))
	_, err = conn.Do("EXPIRE", token, int64(timeLeft))
	if err != nil {
		return err
	}
	return nil
}

