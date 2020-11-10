package Controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Watson-Sei/watson-sei-official/api_v1/Config"
	"github.com/Watson-Sei/watson-sei-official/api_v1/Models"
	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// ユーザー登録
func SignupPost(context *gin.Context)  {
	var user Models.User
	if err := context.Bind(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err":err})
		return
	} else {
		err := Models.CreateUser(user.Username, user.Password)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"err":err})
			return
		} else {
			context.JSON(http.StatusOK, gin.H{"message":"signup success"})
		}
	}
}

// ログイン関数
func LoginPost(context *gin.Context)  {
	var user Models.User
	if err := context.Bind(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err":err})
		return
	} else {
		db := Config.DbConnect()
		loginUser := user.Username
		loginPassword := user.Password
		db.Find(&Models.User{}, "username =?", loginUser).Scan(&user)
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginPassword)); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"err":err, "login": loginUser})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"token": Models.CreateJWTToken(user.Username),
		})
	}
}

// Logout
type Logout struct {
	token	string	`json:"token" biding:"required"`
}

func LogoutPost(context *gin.Context)  {
	token, err := request.ParseFromRequest(context.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		b := []byte(Config.SECRETKEY)
		return b, nil
	})

	if err == nil {
		claims := token.Claims.(jwt2.MapClaims)
		var jwt Logout
		if err = context.Bind(&jwt); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"err":err})
			context.Abort()
		} else {
			c := Config.RedisConnection()
			defer c.Close()
			_, err = c.Do("HSET", jwt.token, "time", claims["exp"])
			_, err = c.Do("SADD", "black-list", jwt.token)
			if err != nil {
				log.Println(err)
			}
			context.JSON(http.StatusOK, gin.H{"message": "ok"})
		}
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"err":err})
	}
}