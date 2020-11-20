package Models

import (
	"time"

	"github.com/Watson-Sei/watson-sei-official/api_v1/Config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username string, password string) (err error) {
	db := Config.DBConnect()
	defer db.Close()
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	// insert処理
	if err := db.Create(&User{Username: username,Password: string(hash)}).Error; err != nil {
		return err
	}
	return nil
}

func PasswordChecker(username string, password string) (err error) {
	var user User
	db := Config.DBConnect()
	defer db.Close()
	db.Find(&User{}, "username =?", username).Scan(&user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}

func CreateJWTToken(username string, userId uint) (map[string]string, error) {
	/*
	アルゴリズム指定
	 */
	// header
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	// claims
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "AccessToken"
	claims["userId"] = userId
	claims["username"] = username
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 4).Unix()

	// Electronic signature
	t, err := token.SignedString([]byte(Config.ACCESS_TOKEN_SECRETKEY))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = "RefreshToken"
	rtClaims["userId"] = userId
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	rt, err := refreshToken.SignedString([]byte(Config.REFRESH_TOKEN_SECRETKEY))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token": t,
		"refresh_token": rt,
	}, nil
}

func RefreshJWTToken(userId uint, token string, exp int64) (map[string]string, error) {
	var user User
	db := Config.DBConnect()
	defer db.Close()
	db.First(&user, userId).Scan(&user)
	// Create New JWT Token (Access, Refresh)
	tokens, err := CreateJWTToken(user.Username, userId)
	if err != nil {
		return nil, err
	}
	// リフレッシュトークンをブラックリストに登録
	err = BlackListSet(exp, token)
	if err != nil {
		return nil, err
	}
	return tokens, nil
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


// BlackListChecker
func BlackListChecker(token string) error {
	conn := Config.RedisConnection()
	defer conn.Close()
	_, err := redis.String(conn.Do("GET", token))
	if err != nil {
		return err
	}
	return nil
}
